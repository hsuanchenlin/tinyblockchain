package types

import "math/big"

type Transaction struct {
	Hash Hash
	From Address
	To Address
	Value big.Int
}


