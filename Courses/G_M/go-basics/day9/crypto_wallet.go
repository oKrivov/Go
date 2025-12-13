package main

type CryptoWallet struct {
	Balance float64
}

func (c *CryptoWallet) Deposit(amount float64) error {
	c.Balance += float64(amount)
	return nil
}

func (c *CryptoWallet) Withdraw(amount float64) error {
	if float64(amount) > c.Balance {
		return ErrInsufficientFunds
	}
	c.Balance -= float64(amount)
	return nil
}

func (c CryptoWallet) GetBalance() float64 {
	return c.Balance
}
