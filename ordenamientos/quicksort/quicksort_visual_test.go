package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestAnimarQuicksortConRepetidos(t *testing.T) {
	array := []int{3, 1, 3, 2, 3}
	original := append([]int(nil), array...)

	var buf bytes.Buffer
	animarQuicksort(&buf, array, 0)

	if !iguales(array, original) {
		t.Errorf("la animación no debe modificar el arreglo original: quedó %v", array)
	}
	if !strings.Contains(buf.String(), "[1 2 3 3 3]") {
		t.Errorf("la animación debe terminar con el arreglo ordenado; salida:\n%s", buf.String())
	}
}

func TestAnimarQuicksort(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	original := append([]int(nil), array...)

	var buf bytes.Buffer
	animarQuicksort(&buf, array, 0)

	salida := buf.String()
	if !iguales(array, original) {
		t.Errorf("la animación no debe modificar el arreglo original: quedó %v", array)
	}
	if !strings.Contains(salida, "Estado inicial") {
		t.Errorf("la animación debe arrancar con el estado inicial; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "[1 2 3 4 5 6 7 8 9 10]") {
		t.Errorf("la animación debe terminar con el arreglo ordenado; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "Particiona") {
		t.Errorf("la animación debe rotular cada partición; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "*") {
		t.Errorf("la animación debe marcar la posición del pivote; salida:\n%s", salida)
	}
}
