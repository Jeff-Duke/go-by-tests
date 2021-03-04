package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds will be displayed when trying to withdraw more than the balance of the wallet
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Crypto is a type alias for our Cryptocurrency
type Crypto int

func (c Crypto) String() string {
	return fmt.Sprintf("%d BTC", c)
}

// Stringer does some stringing
type Stringer interface {
	String() string
}

// Wallet holds coin
type Wallet struct {
	balance Crypto
}

// Balance returns the amount of coin in the wallet
func (w *Wallet) Balance() Crypto {
	return w.balance
}

// Deposit adds new coin to the wallet
func (w *Wallet) Deposit(amt Crypto) {
	w.balance += amt
}

// Withdraw removes coin from the wallet
func (w *Wallet) Withdraw(amt Crypto) error {
	if amt > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amt
	return nil
}
