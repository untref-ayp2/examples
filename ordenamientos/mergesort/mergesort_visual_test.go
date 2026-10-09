package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestAnimarMergesort(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	original := append([]int(nil), array...)

	var buf bytes.Buffer
	animarMergesort(&buf, array, 0)

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
	if !strings.Contains(salida, "Fusiona") {
		t.Errorf("la animación debe rotular cada fusión; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "*") {
		t.Errorf("la animación debe marcar el segmento fusionado; salida:\n%s", salida)
	}
}
