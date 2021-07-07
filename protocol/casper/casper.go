package casper

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/bytom/bytom/common"
	"github.com/bytom/bytom/consensus"
	"github.com/bytom/bytom/errors"
	"github.com/bytom/bytom/protocol/bc"
	"github.com/bytom/bytom/protocol/state"
)

var (
	logModule = "casper"

	errPubKeyIsNotValidator     = errors.New("pub key is not in validators of target checkpoint")
	errVoteToGrowingCheckpoint  = errors.New("validator publish vote to growing checkpoint")
	errVoteToSameCheckpoint     = errors.New("source height and target height in verification is equals")
	errSameHeightInVerification = errors.New("validator publish two distinct votes for the same target height")
	errSpanHeightInVerification = errors.New("validator publish vote within the span of its other votes")
)

// RollbackMsg sent the rollback msg to chain core
type RollbackMsg struct {
	BestHash bc.Hash
	Reply    chan error
}

// Casper is BFT based proof of stack consensus algorithm, it provides safety and liveness in theory,
// it's design mainly refers to https://github.com/ethereum/research/blob/master/papers/casper-basics/casper_basics.pdf
type Casper struct {
	mu    sync.RWMutex
	store state.Store
	tree  *treeNode

	// block hash -> previous checkpoint hash
	prevCheckpointCache *common.Cache
	// block hash + pubKey -> verification
	verificationCache *common.Cache

	rollbackCh chan *RollbackMsg
	newEpochCh chan bc.Hash
}

// NewCasper create a new instance of Casper
// argument checkpoints load the checkpoints from leveldb
// the first element of checkpoints must genesis checkpoint or the last finalized checkpoint in order to reduce memory space
// the others must be successors of first one
func NewCasper(store state.Store, checkpoints []*state.Checkpoint, rollbackCh chan *RollbackMsg) *Casper {
	if checkpoints[0].Height != 0 && checkpoints[0].Status != state.Finalized {
		log.WithFields(log.Fields{"module": logModule}).Panic("first element of checkpoints must genesis or in finalized status")
	}

	casper := &Casper{
		store:               store,
		tree:                makeTree(checkpoints[0], checkpoints[1:]),
		prevCheckpointCache: common.NewCache(1024),
		verificationCache:   common.NewCache(1024),
		rollbackCh:          rollbackCh,
		newEpochCh:          make(chan bc.Hash),
	}
	go casper.authVerificationLoop()
	return casper
}

// LastFinalized return the block height and block hash which is finalized at last
func (c *Casper) LastFinalized() (uint64, bc.Hash) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	root := c.tree.Checkpoint
	return root.Height, root.Hash
}

// LastJustified return the block height and block hash which is justified at last
func (c *Casper) LastJustified() (uint64, bc.Hash) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	node := c.tree.lastJustified()
	return node.Height, node.Hash
}

// Validators return the validators by specified block hash
// e.g. if the block num of epoch is 100, and the block height corresponding to the block hash is 130, then will return the voting results of height in 0~100
func (c *Casper) validators(blockHash *bc.Hash) (map[string]*state.Validator, error) {
	checkpoint, err := c.ParentCheckpoint(blockHash)
	if err != nil {
		return nil, err
	}

	return checkpoint.EffectiveValidators(), nil
}

func (c *Casper) ParentCheckpoint(blockHash *bc.Hash) (*state.Checkpoint, error) {
	hash, err := c.parentCheckpointHash(blockHash)
	if err != nil {
		return nil, err
	}

	return c.store.GetCheckpoint(hash)
}

func (c *Casper) ParentCheckpointByPrevHash(prevBlockHash *bc.Hash) (*state.Checkpoint, error) {
	hash, err := c.parentCheckpointHashByPrevHash(prevBlockHash)
	if err != nil {
		return nil, err
	}

	return c.store.GetCheckpoint(hash)
}

func (c *Casper) bestChain() bc.Hash {
	// root is init justified
	bestNode, _ := c.tree.bestNode(c.tree.Height)
	return bestNode.Hash
}

func (c *Casper) parentCheckpointHash(blockHash *bc.Hash) (*bc.Hash, error) {
	if data, ok := c.prevCheckpointCache.Get(*blockHash); ok {
		return data.(*bc.Hash), nil
	}

	block, err := c.store.GetBlockHeader(blockHash)
	if err != nil {
		return nil, err
	}

	result, err := c.parentCheckpointHashByPrevHash(&block.PreviousBlockHash)
	if err != nil {
		return nil, err
	}

	c.prevCheckpointCache.Add(*blockHash, result)
	return result, nil
}

func (c *Casper) parentCheckpointHashByPrevHash(prevBlockHash *bc.Hash) (*bc.Hash, error) {
	prevHash := prevBlockHash
	for {
		prevBlock, err := c.store.GetBlockHeader(prevHash)
		if err != nil {
			return nil, err
		}

		if prevBlock.Height%consensus.ActiveNetParams.BlocksOfEpoch == 0 {
			c.prevCheckpointCache.Add(*prevBlockHash, prevHash)
			return prevHash, nil
		}

		if data, ok := c.prevCheckpointCache.Get(*prevHash); ok {
			c.prevCheckpointCache.Add(*prevBlockHash, data)
			return data.(*bc.Hash), nil
		}

		prevHash = &prevBlock.PreviousBlockHash
	}
}
