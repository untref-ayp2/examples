package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNuevoEstadoHanoi(t *testing.T) {
	estado := nuevoEstadoHanoi(3)

	esperado := []int{3, 2, 1}
	if len(estado.varas[0]) != len(esperado) {
		t.Fatalf("vara A tiene %d discos; se esperaban %d", len(estado.varas[0]), len(esperado))
	}
	for i, disco := range estado.varas[0] {
		if disco != esperado[i] {
			t.Errorf("vara A, posición %d = %d; se esperaba %d", i, disco, esperado[i])
		}
	}
	if len(estado.varas[1]) != 0 || len(estado.varas[2]) != 0 {
		t.Errorf("las varas B y C deben arrancar vacías")
	}
}

func TestMoverDisco(t *testing.T) {
	estado := nuevoEstadoHanoi(2)

	estado.mover(1, 0, 1)
	estado.mover(2, 0, 2)
	estado.mover(1, 1, 2)

	if len(estado.varas[2]) != 2 {
		t.Fatalf("vara C tiene %d discos; se esperaban 2", len(estado.varas[2]))
	}
	if estado.varas[2][0] != 2 || estado.varas[2][1] != 1 {
		t.Errorf("vara C = %v; se esperaba [2 1]", estado.varas[2])
	}
}

func TestDibujar(t *testing.T) {
	estado := nuevoEstadoHanoi(1)

	var buf bytes.Buffer
	estado.dibujar(&buf, 0)

	salida := buf.String()
	if !strings.Contains(salida, "###") {
		t.Errorf("el dibujo debe contener el disco 1; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "|") {
		t.Errorf("el dibujo debe contener las varas vacías; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "A") || !strings.Contains(salida, "B") || !strings.Contains(salida, "C") {
		t.Errorf("el dibujo debe rotular las varas A, B y C; salida:\n%s", salida)
	}
}
