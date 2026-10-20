package main

import "testing"

func TestCambio(t *testing.T) {
	tests := []struct {
		monto int
		coins []int
		want  int
	}{
		{0, []int{1, 5, 11}, 0},
		{15, []int{1, 5, 11}, 3},
		{22, []int{1, 5, 11}, 2},
		// con solo monedas de 2 no se puede armar un monto impar
		{3, []int{2}, -1},
	}
	for _, tt := range tests {
		if got := Cambio(tt.monto, tt.coins); got != tt.want {
			t.Errorf("Cambio(%d, %v) = %d; want %d", tt.monto, tt.coins, got, tt.want)
		}
	}
}
