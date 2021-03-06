package bc

import "io"

// BlockHeader contains the header information for a blockchain
// block. It satisfies the Entry interface.
func (BlockHeader) typ() string { return "blockheader" }
func (bh *BlockHeader) writeForHash(w io.Writer) {
	mustWriteForHash(w, bh.Version)
	mustWriteForHash(w, bh.Height)
	mustWriteForHash(w, bh.PreviousBlockId)
	mustWriteForHash(w, bh.Timestamp)
	mustWriteForHash(w, bh.TransactionsRoot)
}

// NewBlockHeader creates a new BlockHeader and populates
// its body.
func NewBlockHeader(version, height uint64, previousBlockID *Hash, timestamp uint64, transactionsRoot *Hash) *BlockHeader {
	return &BlockHeader{
		Version:          version,
		Height:           height,
		PreviousBlockId:  previousBlockID,
		Timestamp:        timestamp,
		TransactionsRoot: transactionsRoot,
	}
}
