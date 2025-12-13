package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	ErrInvalidOperator = errors.New("invalid operator")
	ErrDeviaionByZero  = errors.New("division by zero")
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	if err := greetings(reader); err != nil {
		fmt.Println("Error:", err)
		return
	}

	a, b, op, err := getCalcParams(reader)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	res, err := calc(a, b, op)
	if err != nil {
		fmt.Println("Calculation error:", err)
		return
	}
	fmt.Println(res)

	if err := wordsCount(reader); err != nil {
		fmt.Println("Error:", err)
	}
}

func greetings(r *bufio.Reader) error {
	fmt.Println("Enter your name:")

	name, err := readLine(r)
	if err != nil {
		return err
	}

	ageStr, err := readLine(r)
	if err != nil {
		return err
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return fmt.Errorf("invalid age: %w", err)
	}

	fmt.Printf("Hello, %v. Next year your will be %v.\n", name, age+1)
	return nil
}

func readLine(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func getCalcParams(r *bufio.Reader) (int, int, string, error) {
	fmt.Println("Enter first number:")
	aStr, err := readLine(r)
	if err != nil {
		return 0, 0, "", err
	}

	a, err := strconv.Atoi(aStr)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid first numder: %w", err)
	}

	fmt.Println("Enter second number:")

	bStr, err := readLine(r)
	if err != nil {
		return 0, 0, "", err
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid second number: %w", err)
	}

	fmt.Println("Enter operator:")

	op, err := readLine(r)
	if err != nil {
		return 0, 0, "", err
	}

	return a, b, op, nil
}

func calc(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDeviaionByZero
		}
		return a / b, nil
	default:
		return 0, ErrInvalidOperator
	}
}

func wordsCount(r *bufio.Reader) error {
	fmt.Println("Enter text:")
	text, err := readLine(r)
	if err != nil {
		return err
	}

	chars := utf8.RuneCountInString(text)
	words := strings.Fields(text)

	fmt.Println("Words:", len(words))
	fmt.Println("Chars:", chars)

	return nil
}
