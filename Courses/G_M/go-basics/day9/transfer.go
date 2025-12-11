package main

func Transfer(from FinancialAccount, to FinancialAccount, amount float64) error {
	if err := from.Withdraw(amount); err != nil {
		return err
	}

	if err := to.Deposit(amount); err != nil {
		from.Deposit(amount)
		return err
	}

	return nil
}
