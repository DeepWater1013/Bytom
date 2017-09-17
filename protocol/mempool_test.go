package protocol

import (
	"testing"

	"github.com/bytom/protocol/bc/legacy"
	"github.com/bytom/protocol/validation"
)

func TestTxPool(t *testing.T) {
	p := NewTxPool()

	txA := mockCoinbaseTx(1000, 6543)
	txB := mockCoinbaseTx(2000, 2324)
	txC := mockCoinbaseTx(3000, 9322)

	p.AddTransaction(txA, 1, 1000, 5000000000)
	if !p.IsTransactionInPool(&txA.ID) {
		t.Errorf("fail to find added txA in tx pool")
	} else {
		i, _ := p.GetTransaction(&txA.ID)
		if i.Height != 1000 || i.Fee != 5000000000 || i.FeePerKB != 5000000000 {
			t.Errorf("incorrect data of TxDesc structure")
		}
	}

	if p.IsTransactionInPool(&txB.ID) {
		t.Errorf("shouldn't find txB in tx pool")
	}
	p.AddTransaction(txB, 1000, 1, 5000000000)
	if !p.IsTransactionInPool(&txB.ID) {
		t.Errorf("shouldn find txB in tx pool")
	}

	if p.Count() != 2 {
		t.Errorf("get wrong number of tx in the pool")
	}
	p.RemoveTransaction(&txB.ID)
	if p.IsTransactionInPool(&txB.ID) {
		t.Errorf("shouldn't find txB in tx pool")
	}

	p.AddErrCache(&txC.ID)
	if !p.IsTransactionInErrCache(&txC.ID) {
		t.Errorf("shouldn find txC in tx err cache")
	}
	if !p.HaveTransaction(&txC.ID) {
		t.Errorf("shouldn find txC in tx err cache")
	}
}

func mockCoinbaseTx(serializedSize uint64, amount uint64) *legacy.Tx {
	oldTx := &legacy.TxData{
		SerializedSize: serializedSize,
		Outputs: []*legacy.TxOutput{
			legacy.NewTxOutput(*validation.BTMAssetID, amount, []byte{1}, nil),
		},
	}

	return &legacy.Tx{
		TxData: *oldTx,
		Tx:     legacy.MapTx(oldTx),
	}
}
