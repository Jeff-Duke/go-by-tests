package wallet

// Crypto is a type alias for our Cryptocurrency
type Crypto int

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
