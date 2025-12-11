package main

type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.Balance {
		return ErrInsufficientFunds
	}
	a.Balance -= amount
	return nil
}

func (a Account) GetBalance() float64 {
	return float64(a.Balance)
}
