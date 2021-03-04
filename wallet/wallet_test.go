package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Crypto(10))

		assertBalance(t, wallet, Crypto(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Crypto(25)}
		err := wallet.Withdraw(Crypto(10))

		assertBalance(t, wallet, Crypto(15))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Crypto(200)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Crypto(300))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

// Test helpers
func assertBalance(t testing.TB, wallet Wallet, want Crypto) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("caught an error where we didn't expect one")
	}
}
