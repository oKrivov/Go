package main

import (
	"testing"
)

func Test_Deposit(t *testing.T) {
	tests := []struct {
		name string
		// account Account
		balance int
		amount  int
		want    int
	}{
		{
			name:    "first",
			balance: 10,
			amount:  120,
			want:    130,
		},
		{
			name:    "second",
			balance: 20,
			amount:  20,
			want:    40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount(tt.balance)
			account.Deposit(tt.amount)

			if account.Balance != tt.want {
				t.Errorf("Deposit(%v) is not work. The balance(%v), but shout be (%v)", tt.amount, tt.balance, tt.want)
			}
		})

	}
}

func Test_Withdraw_Fail(t *testing.T) {
	tests := []struct {
		name      string
		balance   int
		amount    int
		want      int
		wantErr   bool
		wantError string
	}{
		{
			name:      "fall-withdraw-too-mich",
			balance:   100,
			amount:    120,
			want:      100,
			wantErr:   true,
			wantError: "insufficient funds",
		},
		{
			name:      "success-withdraw",
			balance:   100,
			amount:    20,
			want:      80,
			wantErr:   false,
			wantError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			account := NewAccount(tt.balance)

			err := account.Withdraw(tt.amount)

			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.wantErr && err.Error() != tt.wantError {
				t.Errorf("wrong error message.\n got:  %v\n want: %v", err.Error(), tt.wantError)
			}

			if account.Balance != tt.want {
				t.Errorf("Balance changed! got=%v, want=%v", account.Balance, tt.want)
			}
		})

	}
}
