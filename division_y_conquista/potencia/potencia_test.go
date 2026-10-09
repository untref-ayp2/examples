package main

import "testing"

func TestPotencias(t *testing.T) {
	casos := []struct {
		n, x, esperado int
	}{
		{2, 2, 4},
		{2, 16, 65536},
		{16, 6, 16777216},
	}
	for _, c := range casos {
		got := Potencia(c.n, c.x)
		if got != c.esperado {
			t.Errorf("Potencia(%d, %d) = %d; se esperaba %d", c.n, c.x, got, c.esperado)
		}
	}
}
