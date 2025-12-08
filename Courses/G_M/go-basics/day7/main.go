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

var ErrDivisionByZero = errors.New("division by zero")

type NegativeNumberError struct {
	Value float64
}

func (e NegativeNumberError) Error() string {
	return fmt.Sprintf("cannot take sqrt of negative number: %f", e.Value)
}

// SafeRun выполняет f и безопасно восстанавливает panic.
func SafeRun(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic 🔥:", r)
		}
	}()
	f()
}

// Divide возвращает результат деления a / b или ошибку,
// если b == 0.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("devision by zero")
	}
	return a / b, nil
}

// Sqrt вычисляет квадратный корень, возвращает кастомную ошибку,
// если x < 0.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, NegativeNumberError{Value: x}
	}
	return math.Sqrt(x), nil
}

// readLine читает одну строку из reader и возвращает её trimmed.
// Если EOF или другая ошибка — возвращает её.
func readLine(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

// parseParams парсит строку "a op b" и возвращает a, b, op или ошибку.
// Не использует panic — всегда возвращает ошибку при некорректном вводе.
func parseParams(line string) (int, int, string, error) {
	fields := strings.Fields(line)

	if len(fields) != 3 {
		return 0, 0, "", fmt.Errorf("invalid input format: expected  `a op b`")
	}
	aStr, op, bStr := fields[0], fields[1], fields[2]

	a, err := strconv.Atoi(aStr)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid first numder: %w", err)
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid second numder: %w", err)
	}
	if (op != "+") && op != "-" && op != "*" && op != "/" {
		return 0, 0, "", fmt.Errorf("invalid operator: %s", op)
	}

	return a, b, op, nil
}

// getCalcParams интерактивно запрашивает у пользователя параметры операции.
// Возвращает a,b,op или ошибку (например io.EOF).
func getCalcParams(r *bufio.Reader) (int, int, string, error) {
	for {
		fmt.Println("Enter parmetrs for calculation (example: 2 * 2). Press Ctrl+D to exit.")
		fmt.Print(">")

		line, err := readLine(r)
		if err != nil {
			// распространяется на io.EOF или другие I/O ошибки
			return 0, 0, "", err
		}
		if line == "" {
			fmt.Println("Empty input? try again.")
			continue
		}

		a, b, op, err := parseParams(line)
		if err != nil {
			fmt.Println("Input error:", err)
			continue // попросим ввести снова
		}
		return a, b, op, nil
	}
}

// calc выполняет целочисленную операцию a op b и возвращает результат или ошибку.
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
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
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
