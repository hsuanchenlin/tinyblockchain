package types

import "math/big"

type Transaction struct {
	hash Hash
	from Address
	to Address
	value big.Int
}