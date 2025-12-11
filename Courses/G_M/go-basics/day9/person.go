package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Info() string {
	return fmt.Sprintf("Name = %s, Age = %d", p.Name, p.Age)
}
