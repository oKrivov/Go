package main

import "testing"

func TestNegativeAmountError(t *testing.T) {
	err := NegativeAmountError{Amount: -5}

	if err.Error() == "" {
		t.Errorf("Error() should return non-empty string")
	}
}
