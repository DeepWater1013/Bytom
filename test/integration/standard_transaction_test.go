package integration

import (
	"io/ioutil"
	"os"
	"testing"

	dbm "github.com/tendermint/tmlibs/db"

	"github.com/bytom/blockchain/account"
	"github.com/bytom/blockchain/pseudohsm"
	"github.com/bytom/crypto/ed25519/chainkd"
	"github.com/bytom/protocol/bc/legacy"
	"github.com/bytom/protocol/validation"
	"github.com/bytom/test"
)

func TestP2PKH(t *testing.T) {
	dirPath, err := ioutil.TempDir(".", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dirPath)

	testDB := dbm.NewDB("testdb", "leveldb", "temp")
	defer os.RemoveAll("temp")

	chain, err := test.MockChain(testDB)
	if err != nil {
		t.Fatal(err)
	}

	accountManager := account.NewManager(testDB, chain)
	hsm, err := pseudohsm.New(dirPath)
	if err != nil {
		t.Fatal(err)
	}

	xpub, err := hsm.XCreate("test_pub", "password")
	if err != nil {
		t.Fatal(err)
	}

	testAccount, err := accountManager.Create(nil, []chainkd.XPub{xpub.XPub}, 1, "testAccount", nil)
	if err != nil {
		t.Fatal(err)
	}

	controlProg, err := accountManager.CreateAddress(nil, testAccount.ID, false)
	if err != nil {
		t.Fatal(err)
	}

	utxo := test.MockUTXO(controlProg)
	tpl, tx, err := test.MockTx(utxo, testAccount)
	if err != nil {
		t.Fatal(err)
	}

	if err := test.MockSign(tpl, hsm); err != nil {
		t.Fatal(err)
	}

	if _, _, err = validation.ValidateTx(legacy.MapTx(tx), test.MockBlock()); err != nil {
		t.Fatal(err)
	}
}

func TestP2SH(t *testing.T) {
	dirPath, err := ioutil.TempDir(".", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dirPath)

	testDB := dbm.NewDB("testdb", "leveldb", "temp")
	defer os.RemoveAll("temp")

	chain, err := test.MockChain(testDB)
	if err != nil {
		t.Fatal(err)
	}

	accountManager := account.NewManager(testDB, chain)
	hsm, err := pseudohsm.New(dirPath)
	if err != nil {
		t.Fatal(err)
	}

	xpub1, err := hsm.XCreate("test_pub1", "password")
	if err != nil {
		t.Fatal(err)
	}

	xpub2, err := hsm.XCreate("test_pub2", "password")
	if err != nil {
		t.Fatal(err)
	}

	testAccount, err := accountManager.Create(nil, []chainkd.XPub{xpub1.XPub, xpub2.XPub}, 2, "testAccount", nil)
	if err != nil {
		t.Fatal(err)
	}

	controlProg, err := accountManager.CreateAddress(nil, testAccount.ID, false)
	if err != nil {
		t.Fatal(err)
	}

	utxo := test.MockUTXO(controlProg)
	tpl, tx, err := test.MockTx(utxo, testAccount)
	if err != nil {
		t.Fatal(err)
	}

	if err := test.MockSign(tpl, hsm); err != nil {
		t.Fatal(err)
	}

	if _, _, err = validation.ValidateTx(legacy.MapTx(tx), test.MockBlock()); err != nil {
		t.Fatal(err)
	}
}
