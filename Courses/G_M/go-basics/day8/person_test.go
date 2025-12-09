package main

import "testing"

func TestNewBankUser(t *testing.T) {
	tests := []struct {
		name        string
		age         int
		initBalance int
		want        BankUser
	}{
		{
			name:        "Oleg",
			age:         39,
			initBalance: 10,
			want: BankUser{
				Person:  Person{"Oleg", 39},
				Account: Account{Balance: 10},
			},
		},
		{
			name:        "Kate",
			age:         38,
			initBalance: 111,
			want: BankUser{
				Person:  Person{"Kate", 38},
				Account: Account{Balance: 111},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bUser := NewBankUser(tt.name, tt.age, tt.initBalance)

			if bUser != tt.want {
				t.Errorf("NewBankUser() = %v, want %v", bUser, tt.want)
			}
		})
	}
}
