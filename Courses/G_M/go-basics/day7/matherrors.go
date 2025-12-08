package main

import (
	"errors"
	"fmt"
)

var ErrDivisionByZero = errors.New("division by zero")

type NegativeNumberError struct {
	Value float64
}

func (e NegativeNumberError) Error() string {
	return fmt.Sprintf("cannot take sqrt of negative number: %f", e.Value)
}
