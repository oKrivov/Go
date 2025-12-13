package main

import (
	"fmt"
)

type BankUser struct {
	Person
	Account Account
}

func (b BankUser) Info() {
	fmt.Printf("Name: %v, Age: %d, Balance: %d\n", b.Name, b.Age, b.Account.Balance)
}
