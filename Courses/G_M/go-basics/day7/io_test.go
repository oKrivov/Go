package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestParseParams(t *testing.T) {
	tests := []struct {
		line    string
		a, b    int
		op      string
		wantErr bool
	}{
		{"2 + 3", 2, 3, "+", false},
		{"10 * 5", 10, 5, "*", false},
		{"2 ? 2", 0, 0, "", true},
		{"abc + 2", 0, 0, "", true},
		{"", 0, 0, "", true},
	}

	for _, tt := range tests {
		a, b, op, err := parseParams(tt.line)

		if tt.wantErr && err == nil {
			t.Errorf("expected error for %q", tt.line)
		}

		if !tt.wantErr {
			if a != tt.a || b != tt.b || op != tt.op {
				t.Errorf("wrong parsing for %q", tt.line)
			}
		}

	}
}

func TestReadLine(t *testing.T) {
	input := "hello\n"

	r := bufio.NewReader(strings.NewReader(input))

	line, err := readLine(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if line != "hello" {
		t.Fatalf("expected 'hello', got %q", line)
	}
}
