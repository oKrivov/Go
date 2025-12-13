package main

import "fmt"

var ErrInsufficientFunds = fmt.Errorf("insufficient funds")
var ErrAccountFrozen = fmt.Errorf("account is frozen")

type NegativeAmountError struct {
	Amount int
}

func (e NegativeAmountError) Error() string {
	return fmt.Sprintf("negative amount: %d", e.Amount)
}
