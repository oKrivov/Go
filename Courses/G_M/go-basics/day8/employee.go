package main

import "fmt"

type Employee struct {
	Person
	Position string
	Salary   int
}

func (e Employee) Info() {
	fmt.Printf("Employee: Name=%v, Age=%v, Position=%v, Salary=%v\n",
		e.Name, e.Age, e.Position, e.Salary)
}
