package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestDibujarTablaMochila(t *testing.T) {
	items := []Item{
		{Peso: 3, Valor: 2},
		{Peso: 4, Valor: 3},
		{Peso: 1, Valor: 4},
		{Peso: 2, Valor: 2},
		{Peso: 5, Valor: 5},
	}
	var buf bytes.Buffer
	dibujarTablaMochila(&buf, items, 5, 0)

	salida := buf.String()
	if !strings.Contains(salida, "i=0") || !strings.Contains(salida, "i=5") {
		t.Errorf("la tabla debe tener filas para i=0..len(items); salida:\n%s", salida)
	}
	if !strings.Contains(salida, "Valor máximo: 7") {
		t.Errorf("la tabla debe cerrar con el valor máximo 7; salida:\n%s", salida)
	}
}
