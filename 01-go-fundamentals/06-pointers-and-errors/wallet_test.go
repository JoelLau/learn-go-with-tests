package pointersanderrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	for idx, tt := range []struct {
		wallet  Wallet
		deposit Bitcoin
		expect  Bitcoin
		valid   bool
	}{
		{wallet: Wallet{}, deposit: Bitcoin(10), expect: Bitcoin(10), valid: true},
		{wallet: Wallet{}, deposit: Bitcoin(-10), expect: Bitcoin(0), valid: false},
	} {
		name := fmt.Sprintf("Deposit Test %d", idx)

		t.Run(name, func(t *testing.T) {
			wallet := tt.wallet
			err := wallet.Deposit(tt.deposit)

			assertError(t, err, !tt.valid)
			assertBalance(t, wallet, tt.expect)
		})
	}

	for idx, tt := range []struct {
		wallet   Wallet
		withdraw Bitcoin
		expect   Bitcoin
		valid    bool
	}{
		{wallet: Wallet{10}, withdraw: Bitcoin(10), expect: Bitcoin(0), valid: true},
		{wallet: Wallet{10}, withdraw: Bitcoin(-10), expect: Bitcoin(10), valid: false},
		{wallet: Wallet{10}, withdraw: Bitcoin(100), expect: Bitcoin(10), valid: false},
	} {
		name := fmt.Sprintf("Withdraw Test %d", idx)

		t.Run(name, func(t *testing.T) {
			wallet := tt.wallet
			err := wallet.Withdraw(tt.withdraw)

			assertError(t, err, !tt.valid)
			assertBalance(t, wallet, tt.expect)
		})
	}
}

func assertError(t testing.TB, err error, expectingError bool) {
	t.Helper()

	hasError := err != nil
	have := hasError
	want := expectingError

	if have != want {
		t.Fatalf("have %t want %t", have, want)
	}
}

func assertBalance(t testing.TB, wallet Wallet, balance Bitcoin) {
	t.Helper()

	have := wallet.Balance()
	want := balance

	if have != want {
		t.Errorf("#%v: have %q want %q", wallet, have, want)
	}
}
