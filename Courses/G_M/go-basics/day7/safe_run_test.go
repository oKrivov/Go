package main

import "testing"

func TestSafeRun(t *testing.T) {
	called := false
	SafeRun(func() {
		called = true
	})

	if !called {
		t.Fatalf("function inside SafeRun did not execute")
	}
	SafeRun(func() {
		panic("test")
	})
}
