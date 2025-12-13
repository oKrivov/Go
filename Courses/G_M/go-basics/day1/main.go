package main

import "fmt"

const ARR_LEN = 5

func main() {
	name := "Oleg"
	age := 39
	var arr [ARR_LEN]int
	sum := 0

	fmt.Println("My name is:", name, "\nMy age is:", age, "years old")
	fmt.Println("Result of calling the Add function: ", add(2, 2))

	for i := range arr {
		arr[i] = i + ARR_LEN
		sum += arr[i]
	}
	fmt.Println("Sum of arr: ", sum)
}

func add(a, b int) int {
	return a + b
}
