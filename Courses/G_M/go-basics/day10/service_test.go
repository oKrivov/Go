package main

import (
	"errors"
	"testing"
)

func TestProcessTransaction_InsufficientFunds(t *testing.T) {

	tests := []struct {
		name      string
		balance   int
		amount    int
		targetErr error
	}{
		{
			name:      "success",
			balance:   10,
			amount:    5,
			targetErr: nil,
		},
		{
			name:      "insufficient_funds",
			balance:   10,
			amount:    15,
			targetErr: ErrInsufficientFunds,
		},
		{
			name:    "negative_amount",
			balance: 10,
			amount:  -15,
			targetErr: &NegativeAmountError{
				Amount: -5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{Balance: tt.balance}
			err := ProcessTransaction(&acc, tt.amount)

			if tt.targetErr == nil {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
					return
				}
			}

			if errors.Is(tt.targetErr, ErrInsufficientFunds) {
				if !errors.Is(err, ErrInsufficientFunds) {
					t.Errorf("expected ErrInsufficientFunds got %v", err)
				}
				return
			}
			var negErr *NegativeAmountError
			if errors.As(tt.targetErr, &negErr) {
				if !errors.As(err, &negErr) {
					t.Errorf("expected NegativeAmountError got %v", err)
				}
				return
			}
		})
	}
}
