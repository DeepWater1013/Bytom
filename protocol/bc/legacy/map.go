package legacy

import (
	"github.com/bytom/crypto/sha3pool"
	"github.com/bytom/protocol/bc"
	"github.com/bytom/protocol/vm"
	"github.com/bytom/protocol/vm/vmutil"
)

// MapTx converts a legacy TxData object into its entries-based
// representation.
func MapTx(oldTx *TxData) *bc.Tx {
	txid, header, entries := mapTx(oldTx)

	tx := &bc.Tx{
		TxHeader: header,
		ID:       txid,
		Entries:  entries,
		InputIDs: make([]bc.Hash, len(oldTx.Inputs)),
	}

	var (
		nonceIDs       = make(map[bc.Hash]bool)
		spentOutputIDs = make(map[bc.Hash]bool)
	)
	for id, e := range entries {
		var ord uint64
		switch e := e.(type) {
		case *bc.Issuance:
			anchor, ok := entries[*e.AnchorId]
			if !ok {
				// this tx will be invalid because this issuance is
				// missing an anchor
				continue
			}
			if _, ok := anchor.(*bc.Nonce); ok {
				nonceIDs[*e.AnchorId] = true
			}
			ord = e.Ordinal
			// resume below after the switch

		case *bc.Spend:
			spentOutputIDs[*e.SpentOutputId] = true
			ord = e.Ordinal
			// resume below after the switch

		default:
			continue
		}
		if ord >= uint64(len(oldTx.Inputs)) {
			continue // poorly-formed transaction
		}
		tx.InputIDs[ord] = id
	}

	for id := range nonceIDs {
		tx.NonceIDs = append(tx.NonceIDs, id)
	}
	for id := range spentOutputIDs {
		tx.SpentOutputIDs = append(tx.SpentOutputIDs, id)
	}
	return tx
}

func mapTx(tx *TxData) (headerID bc.Hash, hdr *bc.TxHeader, entryMap map[bc.Hash]bc.Entry) {
	entryMap = make(map[bc.Hash]bc.Entry)

	addEntry := func(e bc.Entry) bc.Hash {
		id := bc.EntryID(e)
		entryMap[id] = e
		return id
	}

	// Loop twice over tx.Inputs, once for spends and once for
	// issuances.  Do spends first so the entry ID of the first spend is
	// available in case an issuance needs it for its anchor.

	var (
		firstSpend   *bc.Spend
		firstSpendID bc.Hash
		spends       []*bc.Spend
		issuances    []*bc.Issuance
		muxSources   = make([]*bc.ValueSource, len(tx.Inputs))
	)

	for i, inp := range tx.Inputs {
		if oldSp, ok := inp.TypedInput.(*SpendInput); ok {
			prog := &bc.Program{VmVersion: oldSp.VMVersion, Code: oldSp.ControlProgram}
			src := &bc.ValueSource{
				Ref:      &oldSp.SourceID,
				Value:    &oldSp.AssetAmount,
				Position: oldSp.SourcePosition,
			}
			out := bc.NewOutput(src, prog, &oldSp.RefDataHash, 0) // ordinal doesn't matter for prevouts, only for result outputs
			prevoutID := addEntry(out)
			refdatahash := hashData(inp.ReferenceData)
			sp := bc.NewSpend(&prevoutID, &refdatahash, uint64(i))
			sp.WitnessArguments = oldSp.Arguments
			id := addEntry(sp)
			muxSources[i] = &bc.ValueSource{
				Ref:   &id,
				Value: &oldSp.AssetAmount,
			}
			if firstSpend == nil {
				firstSpend = sp
				firstSpendID = id
			}
			spends = append(spends, sp)
		}
	}

	for i, inp := range tx.Inputs {
		if oldIss, ok := inp.TypedInput.(*IssuanceInput); ok {
			// Note: asset definitions, initial block ids, and issuance
			// programs are omitted here because they do not contribute to
			// the body hash of an issuance.

			var (
				anchorID    bc.Hash
				setAnchored = func(*bc.Hash) {}
			)

			if len(oldIss.Nonce) > 0 {
				assetID := oldIss.AssetID()

				builder := vmutil.NewBuilder()
				builder.AddData(oldIss.Nonce).AddOp(vm.OP_DROP)
				builder.AddOp(vm.OP_ASSET).AddData(assetID.Bytes()).AddOp(vm.OP_EQUAL)
				prog, _ := builder.Build() // error is impossible

				nonce := bc.NewNonce(&bc.Program{VmVersion: 1, Code: prog})
				anchorID = addEntry(nonce)
				setAnchored = nonce.SetAnchored
			} else if firstSpend != nil {
				anchorID = firstSpendID
				setAnchored = firstSpend.SetAnchored
			}

			val := inp.AssetAmount()

			refdatahash := hashData(inp.ReferenceData)
			assetdefhash := hashData(oldIss.AssetDefinition)
			iss := bc.NewIssuance(&anchorID, &val, &refdatahash, uint64(i))
			iss.WitnessAssetDefinition = &bc.AssetDefinition{
				InitialBlockId: &oldIss.InitialBlock,
				Data:           &assetdefhash,
				IssuanceProgram: &bc.Program{
					VmVersion: oldIss.VMVersion,
					Code:      oldIss.IssuanceProgram,
				},
			}
			iss.WitnessArguments = oldIss.Arguments
			issID := addEntry(iss)
			setAnchored(&issID)

			muxSources[i] = &bc.ValueSource{
				Ref:   &issID,
				Value: &val,
			}
			issuances = append(issuances, iss)
		}
	}

	if tx.IsCoinbase() {
		cb := bc.NewCoinbase()
		cbId := addEntry(cb)

		out := tx.Outputs[0]
		muxSources = []*bc.ValueSource{{
			Ref:   &cbId,
			Value: &out.AssetAmount,
		}}
	}

	mux := bc.NewMux(muxSources, &bc.Program{VmVersion: 1, Code: []byte{byte(vm.OP_TRUE)}})
	muxID := addEntry(mux)

	for _, sp := range spends {
		spentOutput := entryMap[*sp.SpentOutputId].(*bc.Output)
		sp.SetDestination(&muxID, spentOutput.Source.Value, sp.Ordinal)
	}
	for _, iss := range issuances {
		iss.SetDestination(&muxID, iss.Value, iss.Ordinal)
	}

	if tx.IsCoinbase() {
		muxSource := mux.Sources[0]
		cb := entryMap[*muxSource.Ref].(*bc.Coinbase)
		cb.SetDestination(&muxID, muxSource.Value, 0)
	}

	var resultIDs []*bc.Hash

	for i, out := range tx.Outputs {
		src := &bc.ValueSource{
			Ref:      &muxID,
			Value:    &out.AssetAmount,
			Position: uint64(i),
		}
		var dest *bc.ValueDestination
		if vmutil.IsUnspendable(out.ControlProgram) {
			// retirement
			refdatahash := hashData(out.ReferenceData)
			r := bc.NewRetirement(src, &refdatahash, uint64(i))
			rID := addEntry(r)
			resultIDs = append(resultIDs, &rID)
			dest = &bc.ValueDestination{
				Ref:      &rID,
				Position: 0,
			}
		} else {
			// non-retirement
			prog := &bc.Program{out.VMVersion, out.ControlProgram}
			refdatahash := hashData(out.ReferenceData)
			o := bc.NewOutput(src, prog, &refdatahash, uint64(i))
			oID := addEntry(o)
			resultIDs = append(resultIDs, &oID)
			dest = &bc.ValueDestination{
				Ref:      &oID,
				Position: 0,
			}
		}
		dest.Value = src.Value
		mux.WitnessDestinations = append(mux.WitnessDestinations, dest)
	}

	refdatahash := hashData(tx.ReferenceData)
	h := bc.NewTxHeader(tx.Version, resultIDs, &refdatahash, tx.MinTime, tx.MaxTime)
	headerID = addEntry(h)

	return headerID, h, entryMap
}

func mapBlockHeader(old *BlockHeader) (bhID bc.Hash, bh *bc.BlockHeader) {
	bh = bc.NewBlockHeader(old.Version, old.Height, &old.PreviousBlockHash, old.TimestampMS, &old.TransactionsMerkleRoot, &old.AssetsMerkleRoot, old.Nonce, old.Bits)
	bhID = bc.EntryID(bh)
	return
}

func MapBlock(old *Block) *bc.Block {
	if old == nil {
		return nil // if old is nil, so should new be
	}
	b := new(bc.Block)
	b.ID, b.BlockHeader = mapBlockHeader(&old.BlockHeader)
	for _, oldTx := range old.Transactions {
		b.Transactions = append(b.Transactions, oldTx.Tx)
	}
	return b
}

func hashData(data []byte) bc.Hash {
	var b32 [32]byte
	sha3pool.Sum256(b32[:], data)
	return bc.NewHash(b32)
}
