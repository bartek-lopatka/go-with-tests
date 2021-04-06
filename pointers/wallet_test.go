package pointers

import "testing"

func TestWallet(t *testing.T) {

	asssertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.balance
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		asssertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{20}
		wallet.Withdraw(Bitcoin(10))
		asssertBalance(t, wallet, Bitcoin(10))
	})
}
