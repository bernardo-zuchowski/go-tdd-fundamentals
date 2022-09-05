package pointer_errors

import "testing"

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))

		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		w := Wallet{Bitcoin(20)}
		err := w.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		w := Wallet{initialBalance}
		err := w.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, w, initialBalance)
	})
}
