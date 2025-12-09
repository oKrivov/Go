package main

func NewPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func NewEmployee(name string, age int, position string, salary int) Employee {
	return Employee{Person: Person{Name: name, Age: age}, Position: position, Salary: salary}
}

func NewBankUser(name string, age int, initBalance int) BankUser {
	return BankUser{
		Person:  NewPerson(name, age),
		Account: Account{Balance: initBalance},
	}
}

func NewAccount(balance int) *Account {
	return &Account{Balance: balance}
}
