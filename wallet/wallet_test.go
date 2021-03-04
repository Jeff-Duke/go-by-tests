package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Crypto(10))

	got := wallet.Balance()

	want := Crypto(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
