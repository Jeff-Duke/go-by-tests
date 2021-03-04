package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Crypto) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Crypto(10))
		assertBalance(t, wallet, Crypto(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Crypto(25))
		wallet.Withdraw(Crypto(10))
		assertBalance(t, wallet, Crypto(15))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Crypto(200)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Crypto(300))

		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}
