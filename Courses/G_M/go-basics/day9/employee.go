package main

import "fmt"

type Employee struct {
	Person
	Position string
	Salary   int
}

func (e Employee) Info() string {
	return fmt.Sprintf("Name = %s, Age = %d, Position = %s, Salary = %d", e.Name, e.Age, e.Position, e.Salary)
}
