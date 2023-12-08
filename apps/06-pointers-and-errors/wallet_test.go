package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposits", func(t *testing.T) {
	})

	t.Run("withdrawals", func(t *testing.T) {})
}

func assertBalance(t testing.TB, wallet Wallet, balance Bitcoin) {
	t.Helper()

	if wallet.balance != balance {
		t.Errorf("expected balance to be %v, received %v instead", wallet.Balance(), balance)
	}
}

func assertError(t testing.TB, err error, msg string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected %v, but didn't recieve an error", msg)
	}

	if err.Error() != msg {
		t.Fatalf("expected error message '%v', got '%v' instead", err.Error(), msg)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, but received %v", err)
	}
}
