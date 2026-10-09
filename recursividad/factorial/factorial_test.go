package main

import "testing"

func TestFactorialDeCero(t *testing.T) {
	got := Factorial(0)
	if got != 1 {
		t.Errorf("Factorial(0) = %d; se esperaba 1", got)
	}
}

func TestFactorialDe5(t *testing.T) {
	got := Factorial(5)
	if got != 120 {
		t.Errorf("Factorial(5) = %d; se esperaba 120", got)
	}
}
