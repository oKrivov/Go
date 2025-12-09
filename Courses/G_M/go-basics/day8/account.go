package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Balance int
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) error {
	var ErrInsufficientFunds = errors.New("insufficient funds")

	if amount > a.Balance {
		return ErrInsufficientFunds
	}
	a.Balance -= amount
	fmt.Printf("Balance: %v\n", a.Balance)
	return nil
}
