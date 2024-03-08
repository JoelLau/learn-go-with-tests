package pointersanderrors

import "errors"

var ErrNegativeDepositAmount = errors.New("cannot deposit negative amount")
var ErrNegativeWithdrawAmount = errors.New("cannot withdraw negative amount")
var ErrInsufficientBalance = errors.New("insufficient funds")
