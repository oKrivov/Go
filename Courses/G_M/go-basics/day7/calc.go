package main

import (
	"fmt"
	"math"
)

// Divide возвращает результат деления a / b или ошибку,
// если b == 0.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
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
