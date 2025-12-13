package main

import (
	"errors"
	"testing"
)

func Test_Deposit(t *testing.T) {
	tests := []struct {
		name      string
		balance   int
		amount    int
		want      int
		wantErr   bool
		targetErr string
	}{
		{
			name:      "success_deposit",
			balance:   10,
			amount:    100,
			want:      110,
			wantErr:   false,
			targetErr: "",
		},
		{
			name:      "fail_deposit_negative_error",
			balance:   10,
			amount:    -100,
			want:      10,
			wantErr:   true,
			targetErr: "negative amount: -100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := Account{Balance: tt.balance}
			err := account.Deposit(tt.amount)

			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("unepected error: %v", err)
			}

			if tt.wantErr && err.Error() != tt.targetErr {
				t.Errorf("wrong error message.\n got: %v\n want: %v", err.Error(), tt.targetErr)
			}

			if account.Balance != tt.want {
				t.Errorf("Balance changed! got=%v, want=%v", account.Balance, tt.want)
			}
		})
	}
}

func Test_Withdraw(t *testing.T) {
	tests := []struct {
		name      string
		frozen    bool
		amount    int
		balance   int
		want      int
		wantErr   bool
		targetErr error
	}{
		{
			name:    "success",
			frozen:  false,
			amount:  30,
			balance: 100,
			want:    70,
			wantErr: false,
		},
		{
			name:    "negative_amount",
			frozen:  false,
			amount:  -20,
			balance: 100,
			want:    100,
			wantErr: true,
			targetErr: &NegativeAmountError{
				Amount: -20,
			},
		},
		{
			name:      "insufficient_funds",
			frozen:    false,
			amount:    130,
			balance:   100,
			want:      100,
			wantErr:   true,
			targetErr: ErrInsufficientFunds,
		},
		{
			name:      "account_frozen",
			frozen:    true,
			amount:    10,
			balance:   100,
			want:      100,
			wantErr:   true,
			targetErr: ErrAccountFrozen,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			account := &Account{
				Balance: tt.balance,
				Frozen:  tt.frozen,
			}

			err := account.Withdraw(tt.amount)

			if tt.wantErr {
				switch tt.targetErr {
				case ErrInsufficientFunds:
					if !errors.Is(err, ErrInsufficientFunds) {
						t.Errorf("wrong error. got: %v, want: %v", err, ErrInsufficientFunds)
					}
				case ErrAccountFrozen:
					if !errors.Is(err, ErrAccountFrozen) {
						t.Errorf("wrong error. got: %v, want: %v", err, ErrAccountFrozen)
					}

				default:
					var negErr *NegativeAmountError
					if !errors.As(err, &negErr) {
						t.Errorf("wrong error message.\n got : %v\n want: %v", err, tt.targetErr)
					}
				}
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tt.wantErr && account.Balance != tt.want {
				t.Errorf("wrong balance: got %v want %v", account.Balance, tt.want)
			}
		})
	}
}
