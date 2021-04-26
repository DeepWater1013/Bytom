package types

import (
	"io"

	"github.com/bytom/bytom/encoding/blockchain"
	"golang.org/x/crypto/ed25519"
)

// BlockWitness save the consensus node sign
type BlockWitness []byte

func (bw *BlockWitness) readFrom(r *blockchain.Reader) (err error) {
	*bw, err = blockchain.ReadVarstr31(r)
	return err
}

func (bw *BlockWitness) writeTo(w io.Writer) error {
	_, err := blockchain.WriteVarstr31(w, *bw)
	return err
}

func (bw *BlockWitness) Set(data []byte) {
	witness := make([]byte, ed25519.SignatureSize)
	copy(witness, data)
	*bw = witness
}
