package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var ErrDevisionByZero = errors.New("division by zero")

// Exercise 2
type NegativeNumberError struct {
	Value float64
}

func (e NegativeNumberError) Error() string {
	return fmt.Sprintf("cannot take sqrt of negative number: %f", e.Value)
}

func main() {
	// Exercise 1
	Divide(1, 0)
	Divide(10, 3)
	Divide(0, 3)

	// Exercise 2
	Sqrt(-9)
	Sqrt(1)
	Sqrt(2)
	Sqrt(0)

	r := bufio.NewReader(os.Stdin)
	// Exercise 3
	SafeRun(func() {
		fmt.Println("Doing work...")
		panic("something bad happened")
	})
	_, _, txt, err := getCalcParams(r)
	// fmt.Println(err)

	if err != nil {
		fmt.Println("Error;", err)
		return
	}
	fmt.Println(txt)

	// // Exercise 4

	if res, err := calc(r); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(res)
	}
}

// Exercise 1
func Divide(a, b float64) {
	res, err := handlerDivide(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func handlerDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("devision by zero")
	}
	return a / b, nil
}

// Exercise 2
func handlerSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, NegativeNumberError{Value: x}
	}
	return math.Sqrt(x), nil
}

func Sqrt(x float64) {
	res, err := handlerSqrt(x)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(res)
}

// Exercise 3
func SafeRun(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic 🔥:", r)
		}
	}()

	f()
}

// Exercise 4
func readLine(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func parseParams(l string) (int, int, string, error) {
	params := strings.Fields(l)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic 🔥:", r)
		}
	}()

	a, err := strconv.Atoi(params[0])
	if err != nil {
		panic(err)
	}
	op := params[1]

	b, err := strconv.Atoi(params[2])
	if err != nil {
		panic(err)
	}

	return a, b, op, nil
}

func getCalcParams(r *bufio.Reader) (int, int, string, error) {
	var (
		line, op string
		err      error
		a, b     int
	)

	for {
		fmt.Println("Enter parmetrs for calculation.")
		fmt.Println("For example: 2 * 2")
		fmt.Print(": ")

		line, err = readLine(r)

		if err != nil {
			return 0, 0, "", fmt.Errorf("Error: %w", err)
		}

		a, b, op, err = parseParams(line)
		if err != nil {
			return 0, 0, "", fmt.Errorf("Error: %w", err)
		}

		if (op != "+") && op != "-" && op != "*" && op != "/" {
			fmt.Println("You enter incorrect opperrator or not a number!")
			fmt.Println("The operator has been like: + - * /")
		} else if len(line) == 0 {
			fmt.Print("You have not entered the params, please try again!\n")
		} else {
			break
		}
	}

	return a, b, op, err
}

func calc(r *bufio.Reader) (int, error) {
	a, b, op, err := getCalcParams(r)
	if err != nil {
		return 0, err
	}
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDevisionByZero
		}
		return a / b, nil
	default:
		fmt.Println("You enter incorrect opperrator!")
		fmt.Println("The operator has been like: + - * /")
		readLine(r)
		return 0, nil
	}
}
