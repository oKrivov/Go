package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	greetings()

	a, b, op := getCalcParams()
	fmt.Println(calc(a, b, op))

	wordsCount()
}

func greetings() {
	var name string
	var age int

	fmt.Println("Enter your name:")
	_, errName := fmt.Scanln(&name)
	if errName != nil {
		fmt.Println("The value name type is invalid!")
		return
	}

	fmt.Println("Enter your age:")
	_, errAge := fmt.Scanln(&age)
	if errAge != nil {
		fmt.Println("The value age type is invalid!")
		return
	}

	fmt.Printf("Hello, %v. Next year your will be %v.\n", name, age+1)
}

func getCalcParams() (firstNum int, secondNum int, operator string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter first number:")
	text1, _ := reader.ReadString('\n')
	firstNum, err := strconv.Atoi(strings.Trim(text1, "\n"))

	if err == nil {
		fmt.Println("Enter second number:")
		text2, _ := reader.ReadString('\n')
		secondNum, err = strconv.Atoi(strings.Trim(text2, "\n"))

		if err == nil {
			fmt.Println("Enter operator:")
			operator, _ = reader.ReadString('\n')
			operator = strings.Trim(operator, "\n")

			return firstNum, secondNum, operator
		}
	}
	return 0, 0, "0"
}

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

func wordsCount() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text:")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	chars := utf8.RuneCountInString(text)
	words := strings.Fields(text)

	fmt.Println("Words:", len(words))
	fmt.Println("Chars:", chars)
}
