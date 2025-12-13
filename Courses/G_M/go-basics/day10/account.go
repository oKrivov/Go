package main

type Account struct {
	Balance int
	Frozen  bool
}

func (a *Account) Deposit(amount int) error {
	if amount < 0 {
		return &NegativeAmountError{Amount: amount}
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount < 0 {
		return &NegativeAmountError{Amount: amount}
	}

	if a.Frozen {
		return ErrAccountFrozen
	}

	if amount > a.Balance {
		return ErrInsufficientFunds
	}

	a.Balance -= amount
	return nil
}

func (a *Account) Freeze() {
	a.Frozen = true
}
func (a *Account) Unfreeze() {
	a.Frozen = false
}
