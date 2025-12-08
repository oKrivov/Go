package main

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr error
	}{
		{"narmal", 10, 2, 5, nil},
		{"zero numerator", 0, 5, 0, nil},
		{"division by zero", 10, 0, 0, ErrDivisionByZero},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Divide(tt.a, tt.b)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected err %v, got %v", tt.wantErr, err)
			}

			if res != tt.want {
				t.Errorf("expected %v, got %v", tt.want, res)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		name    string
		x       float64
		want    float64
		wantErr bool
	}{
		{"positive", 4, 2, false},
		{"zero", 0, 0, false},
		{"negative", -9, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Sqrt(tt.x)

			if tt.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res != tt.want {
				t.Errorf("expected %v, got %v", tt.want, res)
			}
		})
	}

}
