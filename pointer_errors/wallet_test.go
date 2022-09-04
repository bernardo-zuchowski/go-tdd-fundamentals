package pointer_errors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		w := Wallet{}

		w.Deposit(Bitcoin(10))

		got := w.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("deposit", func(t *testing.T) {
		w := Wallet{Bitcoin(20)}

		w.Withdraw(Bitcoin(10))

		got := w.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
