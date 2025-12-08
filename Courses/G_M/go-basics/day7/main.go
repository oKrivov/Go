package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// SafeRun выполняет f и безопасно восстанавливает panic.
func SafeRun(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic 🔥:", r)
		}
	}()
	f()
}

func main() {

	// Demo SafeRun()
	fmt.Println("\n=== SafeRun Demo ===")
	SafeRun(func() {
		fmt.Println("Doing work...")
		panic("something bad happened")
	})

	// Demo Divide()
	fmt.Println("\n=== Divide Demo ===")
	cases := [][2]float64{
		{1, 0},
		{10, 3},
		{0, 3},
	}

	for _, c := range cases {
		res, err := Divide(c[0], c[1])
		if err != nil {
			if errors.Is(err, ErrDivisionByZero) {
				fmt.Printf("Divide(%v, %v) error: %v\n", c[0], c[1], err)
			} else {
				fmt.Printf("Divide(%v, %v) unexpected error: %v\n", c[0], c[1], err)
			}
			continue
		}
		fmt.Printf("Divide(%v, %v) = %v\n", c[0], c[1], res)
	}

	// Deno Sqrt()
	fmt.Println("\n=== Divide Demo ===")
	sqrtCases := []float64{-9, 1, 2, 0}

	for _, x := range sqrtCases {
		var negError NegativeNumberError
		res, err := Sqrt(x)
		if err != nil {
			if errors.As(err, &negError) {
				fmt.Printf("Sqrt(%v) custom error: %v\n", x, err)
			} else {
				fmt.Printf("Sqrt(%v) unexpected error: %v\n", x, err)
			}
			continue
		}
		fmt.Printf("Sqrt(%v) = %v\n", x, res)
	}

	// Demo CLI Calculator
	fmt.Println("\n=== Interactive calculator ===")
	r := bufio.NewReader(os.Stdin)

	a, b, op, err := getCalcParams(r)

	if err != nil {
		// или io.EOF или пользлватедль набрал Ctrl+D для выхода;
		fmt.Println("Input error (exiting):", err)
		return
	}

	res, err := calc(a, b, op)
	if err != nil {
		if errors.Is(err, ErrDivisionByZero) {
			fmt.Println("Calculation error:", err)
		} else {
			fmt.Println("Unexpected error:", err)
		}
		return
	}
	fmt.Printf("%d %s %d = %d\n", a, op, b, res)
}
