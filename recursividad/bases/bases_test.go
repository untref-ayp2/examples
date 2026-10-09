package main

import "testing"

func TestImprimirEnDecimal(t *testing.T) {
	got := ImprimirEnDecimal(127)
	if got != "127" {
		t.Errorf("ImprimirEnDecimal(127) = %q; se esperaba %q", got, "127")
	}
}

func TestImprimirEnBase(t *testing.T) {
	got := ImprimirEnBase(42, 2)
	if got != "101010" {
		t.Errorf("ImprimirEnBase(42, 2) = %q; se esperaba %q", got, "101010")
	}
}
