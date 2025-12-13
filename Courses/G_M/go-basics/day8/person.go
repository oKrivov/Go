package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Info() {
	fmt.Printf("Person: Name=%v, Age=%v\n", p.Name, p.Age)
}
