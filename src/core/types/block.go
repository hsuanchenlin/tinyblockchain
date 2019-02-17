package types

import (
	"time"
)

type Hash [32]byte
type Address [32]byte

type Header struct {
	Transactions []*Transaction
	BlockHash Hash
	ParentHash Hash
	BlockHeight uint64
	Nonce int
}

type Block struct {
	Header *Header
}

func (blk *Block) addTX(tx *Transaction) (timestamp time.Time, err error) {
	blk.Header.Transactions = append(blk.Header.Transactions, tx)
	return time.Now(), nil
}