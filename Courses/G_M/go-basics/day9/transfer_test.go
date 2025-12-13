package main

import (
	"testing"
)

func Test_Transfer(t *testing.T) {
	tests := []struct {
		name        string
		amount      float64
		balanceFrom float64
		balanceTo   float64
		want        int
		wantErr     bool
	}{
		{
			name:        "succes_transfer",
			amount:      30,
			balanceFrom: 130,
			balanceTo:   0,
			want:        30,
			wantErr:     false,
		},
		{
			name:        "fail_small_balance_to_transer",
			amount:      10,
			balanceFrom: 0,
			balanceTo:   110,
			want:        110,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		from := &Account{tt.balanceFrom}
		to := &Account{tt.balanceTo}

		t.Run(tt.name, func(t *testing.T) {

			err := Transfer(from, to, tt.amount)
			balanceTo := int(to.GetBalance())

			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if balanceTo != tt.want {
				t.Errorf("balance is different! got=%v, want=%v", balanceTo, tt.want)
			}

		})
	}
}

func Test_Account_To_Cryto_Transfer(t *testing.T) {
	tests := []struct {
		name        string
		amount      float64
		balanceFrom float64
		balanceTo   float64
		want        float64
		wantErr     bool
	}{
		{
			name:        "success_account_to_crypto_transfer",
			amount:      11,
			balanceFrom: 13,
			balanceTo:   2.01,
			want:        13.01,
			wantErr:     false,
		},
		{
			name:        "fail_account_to_crypto_transfer",
			amount:      14,
			balanceFrom: 13,
			balanceTo:   2.01,
			want:        2.01,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		from := &Account{tt.balanceFrom}
		to := &CryptoWallet{tt.balanceTo}

		t.Run(tt.name, func(t *testing.T) {

			err := Transfer(from, to, tt.amount)
			balanceTo := to.GetBalance()

			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if balanceTo != tt.want {
				t.Errorf("balance is different! got=%v, want=%v", balanceTo, tt.want)
			}

		})
	}
}

func Test_Crypto_To_Account_Tranfer(t *testing.T) {
	tests := []struct {
		name        string
		amount      float64
		balanceFrom float64
		balanceTo   float64
		want        int
		wantErr     bool
	}{
		{
			name:        "success_crypto_to_account_transfer",
			amount:      11,
			balanceFrom: 13.05,
			balanceTo:   2,
			want:        13,
			wantErr:     false,
		},
		{
			name:        "fail_crypto_to_account_transfer",
			amount:      1,
			balanceFrom: 0.3,
			balanceTo:   12,
			want:        12,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from := &CryptoWallet{tt.balanceFrom}
			to := &Account{tt.balanceTo}

			t.Run(tt.name, func(t *testing.T) {

				err := Transfer(from, to, tt.amount)
				balanceTo := int(to.GetBalance())

				if tt.wantErr && err == nil {
					t.Fatal("expected error, got nil")
				}
				if !tt.wantErr && err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if balanceTo != tt.want {
					t.Errorf("balance is different! got=%v, want=%v", balanceTo, tt.want)
				}

			})

		})
	}
}
