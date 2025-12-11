package main

type Infoer interface {
	Info() string
}

type Logger interface {
	Log(message string)
}

type FinancialAccount interface {
	Deposit(float64) error
	Withdraw(float64) error
	GetBalance() float64
}
