package types

type Hash [32]byte
type Address [20]byte

type Header struct {
	Transactions []*Transaction
	BlockHash Hash
	ParentHash Hash
	BlockHeight uint64
}

type Block struct {
	header *Header
}