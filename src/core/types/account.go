package types

import "math/big"

type Account struct {
	address Address
	amount big.Int
	nonce uint64
}


