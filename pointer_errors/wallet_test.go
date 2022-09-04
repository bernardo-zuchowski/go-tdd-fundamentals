package pointer_errors

import "testing"

var assertBalance = func(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()

	got := w.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))

		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		w := Wallet{10}
		w.Withdraw(Bitcoin(10))

		assertBalance(t, w, Bitcoin(0))
	})
}
