package database

import (
	"os"
	"testing"

	"github.com/bytom/bytom/config"
	dbm "github.com/bytom/bytom/database/leveldb"
	"github.com/bytom/bytom/database/storage"
	"github.com/bytom/bytom/protocol"
	"github.com/bytom/bytom/protocol/bc"
	"github.com/bytom/bytom/protocol/bc/types"
	"github.com/bytom/bytom/protocol/state"
	"github.com/bytom/bytom/testutil"
)

func TestSaveChainStatus(t *testing.T) {
	defer os.RemoveAll("temp")
	testDB := dbm.NewDB("testdb", "leveldb", "temp")
	store := NewStore(testDB)

	blockHeader := &types.BlockHeader{Height: 100}
	blockHash := blockHeader.Hash() //Hash: bc.Hash{V0: 0, V1: 1, V2: 2, V3: 3}
	view := &state.UtxoViewpoint{
		Entries: map[bc.Hash]*storage.UtxoEntry{
			bc.Hash{V0: 1, V1: 2, V2: 3, V3: 4}: &storage.UtxoEntry{IsCoinBase: false, BlockHeight: 100, Spent: false},
			bc.Hash{V0: 1, V1: 2, V2: 3, V3: 4}: &storage.UtxoEntry{IsCoinBase: true, BlockHeight: 100, Spent: true},
			bc.Hash{V0: 1, V1: 1, V2: 3, V3: 4}: &storage.UtxoEntry{IsCoinBase: false, BlockHeight: 100, Spent: true},
		},
	}

	contractView := state.NewContractViewpoint()
	if err := store.SaveChainStatus(blockHeader, blockHeader, []*types.BlockHeader{blockHeader}, view, contractView); err != nil {
		t.Fatal(err)
	}

	expectStatus := &protocol.BlockStoreState{Height: blockHeader.Height, Hash: &blockHash, FinalizedHeight: blockHeader.Height, FinalizedHash: &blockHash}
	if !testutil.DeepEqual(store.GetStoreStatus(), expectStatus) {
		t.Errorf("got block status:%v, expect block status:%v", store.GetStoreStatus(), expectStatus)
	}

	for hash, utxo := range view.Entries {
		if utxo.Spent && !utxo.IsCoinBase {
			continue
		}

		gotUtxo, err := store.GetUtxo(&hash)
		if err != nil {
			t.Fatal(err)
		}

		if !testutil.DeepEqual(utxo, gotUtxo) {
			t.Errorf("got utxo entry:%v, expect utxo entry:%v", gotUtxo, utxo)
		}
	}
}

func TestSaveBlock(t *testing.T) {
	testDB := dbm.NewDB("testdb", "leveldb", "temp")
	defer func() {
		testDB.Close()
		os.RemoveAll("temp")
	}()

	store := NewStore(testDB)
	block := config.GenesisBlock()
	if err := store.SaveBlock(block); err != nil {
		t.Fatal(err)
	}

	blockHash := block.Hash()
	gotBlock, err := store.GetBlock(&blockHash)
	if err != nil {
		t.Fatal(err)
	}

	gotBlock.Transactions[0].Tx.SerializedSize = 0
	gotBlock.Transactions[0].SerializedSize = 0
	if !testutil.DeepEqual(block, gotBlock) {
		t.Errorf("got block:%v, expect block:%v", gotBlock, block)
	}

	data := store.db.Get(CalcBlockHeaderIndexKey(block.Height, &blockHash))
	gotBlockHeader := types.BlockHeader{}
	if err := gotBlockHeader.UnmarshalText(data); err != nil {
		t.Fatal(err)
	}

	if !testutil.DeepEqual(block.BlockHeader, gotBlockHeader) {
		t.Errorf("got block header:%v, expect block header:%v", gotBlockHeader, block.BlockHeader)
	}
}
