package main

import "fmt"

func main() {

	p := NewPerson("Oleg", 39)
	e := NewEmployee(p.Name, p.Age, "sturdent", 500)
	bUser := NewBankUser(p.Name, p.Age, 0)
	p.Info()
	e.Info()
	bUser.Info()

	bUser.Account.Deposit(130)

	bUser.Account.Deposit(240)

	if err := bUser.Account.Withdraw(300); err != nil {
		fmt.Println(err)
	}

	if err := bUser.Account.Withdraw(1000); err != nil {
		fmt.Println(err)
	}
}
