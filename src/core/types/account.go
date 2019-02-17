package types

type Account struct {
	address Address
	amount  int
	nonce uint64
}


//type AccountTree struct {
//	root
//}

func (acc *Account) init(balance int) {
	acc.amount = balance
}

func (acc *Account) send(amount int) {
	acc.amount -= amount
}

func (acc *Account) receive(amount int) {
	acc.amount += amount
}