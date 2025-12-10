package main

import (
	"errors"
	"fmt"
)

// -------------------------
// Интерфейсы
// -------------------------

type Infoer interface {
	Info() string
}

type AccountOps interface {
	Deposit(int)
	Withdraw(int) error
}

var ErrInsufficientFunds = errors.New("insufficient funds")

// -------------------------
// Структуры и методы
// -------------------------

type Person struct {
	Name string
	Age  int
}

func (p Person) Info() string {
	return fmt.Sprintf("Name = %s, Age = %d", p.Name, p.Age)
}

type Employee struct {
	Person
	Position string
	Salary   int
}

func (e Employee) Info() string {
	return fmt.Sprintf("Name = %s, Age = %d, Position = %s, Salary = %d", e.Name, e.Age, e.Position, e.Salary)
}

type BankUser struct {
	Person
	Account
}

func (b BankUser) Info() string {
	return fmt.Sprintf("Name = %s, Age = %d, Balance = %d", b.Person.Name, b.Person.Age, b.Account.Balance)
}

type Account struct {
	Balance int
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if amount > a.Balance {
		return ErrInsufficientFunds
	}
	a.Balance -= amount

	return nil
}

type CryptoWallet struct {
	Balance float64
}

func (c *CryptoWallet) Deposit(amount int) {
	c.Balance += float64(amount)
}

func (c *CryptoWallet) Withdraw(amount int) error {
	if float64(amount) > c.Balance {
		return ErrInsufficientFunds
	}
	c.Balance -= float64(amount)

	return nil
}

// -------------------------
// Полиморфные функции
// -------------------------

func Pay(a AccountOps, amount int) error {
	err := a.Withdraw(amount)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	fmt.Println("Payment successful")
	return nil
}

func PrintInfo(i Infoer) {
	fmt.Println(i.Info())
}

// -------------------------
// main
// -------------------------

func main() {
	p := Person{"Oleg", 39}
	e := Employee{p, "Student", 0}
	a := Account{4410}
	bUser := BankUser{p, a}

	PrintInfo(p)
	PrintInfo(e)
	PrintInfo(bUser)

	var acc AccountOps = &a
	Pay(acc, 300)
}
