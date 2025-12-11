package main

import "fmt"

type BankUser struct {
	Person
	Account
}

func (b BankUser) Info() string {
	return fmt.Sprintf("Name = %s, Age = %d, Balance = %v", b.Person.Name, b.Person.Age, b.Account.Balance)
}
