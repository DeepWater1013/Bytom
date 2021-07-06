package protocol

import (
	"encoding/hex"

	log "github.com/sirupsen/logrus"

	"github.com/bytom/bytom/crypto/ed25519/chainkd"
	"github.com/bytom/bytom/errors"
	"github.com/bytom/bytom/event"
	"github.com/bytom/bytom/protocol/bc"
	"github.com/bytom/bytom/protocol/bc/types"
	"github.com/bytom/bytom/protocol/casper"
	"github.com/bytom/bytom/protocol/state"
	"github.com/bytom/bytom/protocol/validation"
)

var (
	// ErrBadBlock is returned when a block is invalid.
	ErrBadBlock = errors.New("invalid block")
	// ErrBadStateRoot is returned when the computed assets merkle root
	// disagrees with the one declared in a block header.
	ErrBadStateRoot = errors.New("invalid state merkle root")
)

// BlockExist check is a block in chain or orphan
func (c *Chain) BlockExist(hash *bc.Hash) bool {
	if _, err := c.store.GetBlockHeader(hash); err == nil {
		return true
	}

	return c.orphanManage.BlockExist(hash)
}

// GetBlockByHash return a block by given hash
func (c *Chain) GetBlockByHash(hash *bc.Hash) (*types.Block, error) {
	return c.store.GetBlock(hash)
}

// GetBlockByHeight return a block header by given height
func (c *Chain) GetBlockByHeight(height uint64) (*types.Block, error) {
	hash, err := c.store.GetMainChainHash(height)
	if err != nil {
		return nil, errors.Wrap(err, "can't find block in given height")
	}

	return c.store.GetBlock(hash)
}

// GetHeaderByHash return a block header by given hash
func (c *Chain) GetHeaderByHash(hash *bc.Hash) (*types.BlockHeader, error) {
	return c.store.GetBlockHeader(hash)
}

// GetHeaderByHeight return a block header by given height
func (c *Chain) GetHeaderByHeight(height uint64) (*types.BlockHeader, error) {
	hash, err := c.store.GetMainChainHash(height)
	if err != nil {
		return nil, errors.Wrap(err, "can't find block header in given height")
	}

	return c.store.GetBlockHeader(hash)
}

func (c *Chain) calcReorganizeChain(beginAttach *types.BlockHeader, beginDetach *types.BlockHeader) ([]*types.BlockHeader, []*types.BlockHeader, error) {
	var err error
	var attachBlockHeaders []*types.BlockHeader
	var detachBlockHeaders []*types.BlockHeader

	for attachBlockHeader, detachBlockHeader := beginAttach, beginDetach; detachBlockHeader.Hash() != attachBlockHeader.Hash(); {
		var attachRollback, detachRollBack bool
		if attachRollback = attachBlockHeader.Height >= detachBlockHeader.Height; attachRollback {
			attachBlockHeaders = append([]*types.BlockHeader{attachBlockHeader}, attachBlockHeaders...)
		}

		if detachRollBack = attachBlockHeader.Height <= detachBlockHeader.Height; detachRollBack {
			detachBlockHeaders = append(detachBlockHeaders, detachBlockHeader)
		}

		if attachRollback {
			attachBlockHeader, err = c.store.GetBlockHeader(&attachBlockHeader.PreviousBlockHash)
			if err != nil {
				return nil, nil, err
			}
		}

		if detachRollBack {
			detachBlockHeader, err = c.store.GetBlockHeader(&detachBlockHeader.PreviousBlockHash)
			if err != nil {
				return nil, nil, err
			}
		}
	}
	return attachBlockHeaders, detachBlockHeaders, nil
}

func (c *Chain) connectBlock(block *types.Block) (err error) {
	bcBlock := types.MapBlock(block)
	utxoView := state.NewUtxoViewpoint()
	if err := c.store.GetTransactionsUtxo(utxoView, bcBlock.Transactions); err != nil {
		return err
	}
	if err := utxoView.ApplyBlock(bcBlock); err != nil {
		return err
	}

	reply, err := c.casper.ApplyBlock(block)
	if err != nil {
		return err
	}

	if reply.Verification != nil {
		if err := c.broadcastVerification(reply.Verification); err != nil {
			return err
		}
	}

	contractView := state.NewContractViewpoint()
	contractView.ApplyBlock(block)

	if err := c.setState(&block.BlockHeader, []*types.BlockHeader{&block.BlockHeader}, utxoView, contractView); err != nil {
		return err
	}

	for _, tx := range block.Transactions {
		c.txPool.RemoveTransaction(&tx.Tx.ID)
	}
	return nil
}

func (c *Chain) reorganizeChain(blockHeader *types.BlockHeader) error {
	attachNodes, detachNodes, err := c.calcReorganizeChain(blockHeader, c.bestBlockHeader)
	if err != nil {
		return err
	}

	utxoView := state.NewUtxoViewpoint()
	contractView := state.NewContractViewpoint()

	txsToRestore := map[bc.Hash]*types.Tx{}
	for _, detachNode := range detachNodes {
		hash := detachNode.Hash()
		b, err := c.store.GetBlock(&hash)
		if err != nil {
			return err
		}

		detachBlock := types.MapBlock(b)
		if err := c.store.GetTransactionsUtxo(utxoView, detachBlock.Transactions); err != nil {
			return err
		}

		if err := utxoView.DetachBlock(detachBlock); err != nil {
			return err
		}

		contractView.DetachBlock(b)
		for _, tx := range b.Transactions {
			txsToRestore[tx.ID] = tx
		}
		log.WithFields(log.Fields{"module": logModule, "height": detachNode.Height, "hash": hash.String()}).Debug("detach from mainchain")
	}

	txsToRemove := map[bc.Hash]*types.Tx{}
	for _, attachNode := range attachNodes {
		hash := attachNode.Hash()
		b, err := c.store.GetBlock(&hash)
		if err != nil {
			return err
		}

		attachBlock := types.MapBlock(b)
		if err := c.store.GetTransactionsUtxo(utxoView, attachBlock.Transactions); err != nil {
			return err
		}

		if err := utxoView.ApplyBlock(attachBlock); err != nil {
			return err
		}

		contractView.ApplyBlock(b)
		for _, tx := range b.Transactions {
			if _, ok := txsToRestore[tx.ID]; !ok {
				txsToRemove[tx.ID] = tx
			} else {
				delete(txsToRestore, tx.ID)
			}
		}

		log.WithFields(log.Fields{"module": logModule, "height": attachNode.Height, "hash": hash.String()}).Debug("attach from mainchain")
	}

	if err := c.setState(blockHeader, []*types.BlockHeader{blockHeader}, utxoView, contractView); err != nil {
		return err
	}

	for txHash := range txsToRemove {
		c.txPool.RemoveTransaction(&txHash)
	}

	for _, tx := range txsToRestore {
		// the number of restored Tx should be very small or most of time ZERO
		// Error returned from validation is ignored, tx could still be lost if validation fails.
		// TODO: adjust tx timestamp so that it won't starve in pool.
		if _, err := c.ValidateTx(tx); err != nil {
			log.WithFields(log.Fields{"module": logModule, "tx_id": tx.Tx.ID.String(), "error": err}).Info("restore tx fail")
		}
	}

	if len(txsToRestore) > 0 {
		log.WithFields(log.Fields{"module": logModule, "num": len(txsToRestore)}).Debug("restore txs back to pool")
	}

	return nil
}

func (c *Chain) broadcastVerification(v *casper.Verification) error {
	pubKey, err := hex.DecodeString(v.PubKey)
	if err != nil {
		return err
	}

	return c.eventDispatcher.Post(event.BlockVerificationEvent{
		SourceHeight: v.SourceHeight,
		SourceHash:   v.SourceHash,
		TargetHeight: v.TargetHeight,
		TargetHash:   v.TargetHash,
		PubKey:       pubKey,
		Signature:    v.Signature,
	})
}

// SaveBlock will validate and save block into storage
func (c *Chain) saveBlock(block *types.Block) error {
	if err := c.verifyBlockSignature(&block.BlockHeader); err != nil {
		return err
	}

	bcBlock := types.MapBlock(block)
	parent, err := c.store.GetBlockHeader(&block.PreviousBlockHash)
	if err != nil {
		return err
	}

	checkpoint, err := c.PrevCheckpointByPrevHash(&block.PreviousBlockHash)
	if err != nil {
		return err
	}

	if err := validation.ValidateBlock(bcBlock, parent, checkpoint, c.ProgramConverter); err != nil {
		return errors.Sub(ErrBadBlock, err)
	}

	if err := c.store.SaveBlock(block); err != nil {
		return err
	}

	c.orphanManage.Delete(&bcBlock.ID)
	return nil
}

func (c *Chain) saveSubBlock(block *types.Block) *types.Block {
	blockHash := block.Hash()
	prevOrphans, ok := c.orphanManage.GetPrevOrphans(&blockHash)
	if !ok {
		return block
	}

	bestBlock := block
	for _, prevOrphan := range prevOrphans {
		orphanBlock, ok := c.orphanManage.Get(prevOrphan)
		if !ok {
			log.WithFields(log.Fields{"module": logModule, "hash": prevOrphan.String()}).Warning("saveSubBlock fail to get block from orphanManage")
			continue
		}
		if err := c.saveBlock(orphanBlock); err != nil {
			log.WithFields(log.Fields{"module": logModule, "hash": prevOrphan.String(), "height": orphanBlock.Height}).Warning("saveSubBlock fail to save block")
			continue
		}

		if subBestBlock := c.saveSubBlock(orphanBlock); subBestBlock.Height > bestBlock.Height {
			bestBlock = subBestBlock
		}
	}
	return bestBlock
}

func (c *Chain) verifyBlockSignature(blockHeader *types.BlockHeader) error {
	validator, err := c.GetValidator(&blockHeader.PreviousBlockHash, blockHeader.Timestamp)
	if err != nil {
		return err
	}

	xPub := chainkd.XPub{}
	pubKey, _ := hex.DecodeString(validator.PubKey)
	copy(xPub[:], pubKey)
	if ok := xPub.Verify(blockHeader.Hash().Bytes(), blockHeader.BlockWitness); !ok {
		return errors.New("fail to verify block header signature")
	}

	return nil
}

type processBlockResponse struct {
	isOrphan bool
	err      error
}

type processBlockMsg struct {
	block *types.Block
	reply chan processBlockResponse
}

// ProcessBlock is the entry for chain update
func (c *Chain) ProcessBlock(block *types.Block) (bool, error) {
	reply := make(chan processBlockResponse, 1)
	c.processBlockCh <- &processBlockMsg{block: block, reply: reply}
	response := <-reply
	return response.isOrphan, response.err
}

func (c *Chain) blockProcessor() {
	for {
		select {
		case msg := <-c.processBlockCh:
			isOrphan, err := c.processBlock(msg.block)
			msg.reply <- processBlockResponse{isOrphan: isOrphan, err: err}
		case msg := <-c.processRollbackCh:
			err := c.rollback(msg.BestHash)
			msg.Reply <- err
		}
	}
}

// ProcessBlock is the entry for handle block insert
func (c *Chain) processBlock(block *types.Block) (bool, error) {
	blockHash := block.Hash()
	if c.BlockExist(&blockHash) && c.bestBlockHeader.Height >= block.Height {
		log.WithFields(log.Fields{"module": logModule, "hash": blockHash.String(), "height": block.Height}).Info("block has been processed")
		return c.orphanManage.BlockExist(&blockHash), nil
	}

	if _, err := c.store.GetBlockHeader(&block.PreviousBlockHash); err != nil {
		c.orphanManage.Add(block)
		return true, nil
	}

	if err := c.saveBlock(block); err != nil {
		return false, err
	}

	bestBlock := c.saveSubBlock(block)
	bestBlockHeader := &bestBlock.BlockHeader

	if bestBlockHeader.PreviousBlockHash == c.bestBlockHeader.Hash() {
		log.WithFields(log.Fields{"module": logModule}).Debug("append block to the end of mainchain")
		return false, c.connectBlock(bestBlock)
	}

	return false, c.applyForkChainToCasper(bestBlockHeader)
}

func (c *Chain) applyForkChainToCasper(beginAttach *types.BlockHeader) error {
	attachNodes, _, err := c.calcReorganizeChain(beginAttach, c.bestBlockHeader)
	if err != nil {
		return err
	}

	var reply *casper.ApplyBlockReply
	for _, node := range attachNodes {
		hash := node.Hash()
		block, err := c.store.GetBlock(&hash)
		if err != nil {
			return err
		}

		reply, err = c.casper.ApplyBlock(block)
		if err != nil {
			return err
		}

		log.WithFields(log.Fields{"module": logModule, "height": node.Height, "hash": hash.String()}).Info("apply fork node")

		if reply.Verification != nil {
			if err := c.broadcastVerification(reply.Verification); err != nil {
				return err
			}
		}
	}

	if reply.BestHash != c.bestBlockHeader.Hash() {
		return c.rollback(reply.BestHash)
	}

	return nil
}

func (c *Chain) rollback(bestHash bc.Hash) error {
	if c.bestBlockHeader.Hash() == bestHash {
		return nil
	}

	blockHeader, err := c.GetHeaderByHash(&bestHash)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"module": logModule, "bestHash": bestHash.String()}).Info("start to reorganize chain")
	return c.reorganizeChain(blockHeader)
}
