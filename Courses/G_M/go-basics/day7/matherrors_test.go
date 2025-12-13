package main

import (
	"testing"
)

func TestNegativeNumberError(t *testing.T) {
	err := NegativeNumberError{Value: -5}

	if err.Error() == "" {
		t.Errorf("Error() should return non-empty string")
	}
}
