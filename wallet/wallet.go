package wallet

import "fmt"

// Crypto is a type alias for our Cryptocurrency
type Crypto int

func (c Crypto) String() string {
	return fmt.Sprintf("%d BTC", c)
}

// Wallet holds coin
type Wallet struct {
	balance Crypto
}

// Stringer does some stringing
type Stringer interface {
	String() string
}

// Balance returns the amount of coin in the wallet
func (w *Wallet) Balance() Crypto {
	return w.balance
}

// Deposit adds new coin to the wallet
func (w *Wallet) Deposit(amt Crypto) {
	w.balance += amt
}
