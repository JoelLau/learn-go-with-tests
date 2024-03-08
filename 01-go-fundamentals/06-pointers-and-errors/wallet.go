package pointersanderrors

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amt Bitcoin) error {
	if amt < 0 {
		return ErrNegativeDepositAmount
	}

	w.balance += amt
	return nil
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt < 0 {
		return ErrNegativeWithdrawAmount
	}

	if amt > w.balance {
		return ErrInsufficientBalance
	}

	w.balance -= amt
	return nil
}
