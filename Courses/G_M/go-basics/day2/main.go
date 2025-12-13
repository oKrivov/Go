package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxNum := 5
	num := rand.Intn(10)

	if num > maxNum {
		num = (num * -1)
	}

	checkNumber(num)

	operators := []string{"+", "-", "*", "/"}
	fmt.Println("calc function result: ", calc(num, rand.Intn(10), operators[rand.Intn(4)]))

	age := rand.Intn(100)
	whatAge(age)
}

func checkNumber(n int) {
	if n > 0 {
		fmt.Println("positive")
	} else if n < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}
}

/*
op = "+" → вернуть a + b
op = "-" → вернуть a - b
op = "*" → вернуть a * b
op = "/" → вернуть a / b
*/
func calc(a, b int, op string) int {
	result := 0
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("you can't divide by zero!")
			return 0
		}
		result = a / b
	default:
		fmt.Println("unknown operator:", op)
		return 0
	}
	return result
}

/*
	“child” если < 14

“teenager” если 14–17
“adult” если 18–64
“senior” если > 64
*/
func whatAge(age int) {
	switch {
	case age > 64:
		fmt.Println("senior")
	case age > 17:
		fmt.Println("adult")
	case age >= 14:
		fmt.Println("teenager")
	default:
		fmt.Println("child")
	}
}
