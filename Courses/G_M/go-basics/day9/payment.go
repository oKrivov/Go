package main

import "fmt"

func Pay(fAcc FinancialAccount, amount float64, logger Logger) error {
	logger.Log(fmt.Sprintf("Trying to withdraw %v", amount))

	err := fAcc.Withdraw(amount)
	if err != nil {
		logger.Log(fmt.Sprintf("Error: %v", err))
		return err
	}

	logger.Log("Success")
	return nil
}

func PrintInfo(i Infoer) {
	fmt.Println(i.Info())
}
