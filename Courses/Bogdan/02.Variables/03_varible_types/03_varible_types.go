package main

import "fmt"

func main() {
	name := "Americano"
	price := 2.99
	ready := true
	orderedCount := 5
	var stockCount int64 = 5000

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", price)
	fmt.Printf("%T\n", ready)
	fmt.Printf("%T\n", orderedCount)
	fmt.Printf("%T\n", stockCount)
}
