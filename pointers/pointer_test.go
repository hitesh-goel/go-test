package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	t.Run("Deposit", func(t *testing.T) {
		wallet.Deposit(10)
		want := BitCoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Widraw", func(t *testing.T) {
		err := wallet.Widraw(5)
		want := BitCoin(5)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Widraw more funds", func(t *testing.T) {
		err := wallet.Widraw(20)
		assertBalance(t, wallet, BitCoin(5))
		assertError(t, err)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didnt get one")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error("got an error but didnt want one")
	}
}
