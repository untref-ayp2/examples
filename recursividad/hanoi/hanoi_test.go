package main

import (
	"fmt"
	"strings"
	"testing"
)

// TestMoverConTresDiscos verifica la secuencia exacta de movimientos para la
// torre clásica de 3 discos.
func TestMoverConTresDiscos(t *testing.T) {
	var movimientos []string

	Mover(3, "A", "B", "C", func(disco int, origen, destino string) {
		movimientos = append(movimientos, fmt.Sprintf("%d:%s->%s", disco, origen, destino))
	})

	esperado := "1:A->C,2:A->B,1:C->B,3:A->C,1:B->A,2:B->C,1:A->C"
	got := strings.Join(movimientos, ",")

	if got != esperado {
		t.Errorf("secuencia de movimientos = %s; se esperaba %s", got, esperado)
	}
}

// TestCantidadDeMovimientos verifica que con n discos se generan 2^n - 1
// movimientos.
func TestCantidadDeMovimientos(t *testing.T) {
	contador := 0
	Mover(5, "A", "B", "C", func(int, string, string) { contador++ })

	if contador != 31 {
		t.Errorf("Mover(5, ...) generó %d movimientos; se esperaban 31", contador)
	}
}
