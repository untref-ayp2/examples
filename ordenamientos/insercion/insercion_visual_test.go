package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestAnimarInsercion(t *testing.T) {
	arr := []int{5, 2, 4, 6, 1, 3}
	original := append([]int(nil), arr...)

	var buf bytes.Buffer
	animarInsercion(&buf, arr, 0)

	salida := buf.String()
	if !iguales(arr, original) {
		t.Errorf("la animación no debe modificar el arreglo original: quedó %v", arr)
	}
	if !strings.Contains(salida, "Estado inicial") {
		t.Errorf("la animación debe arrancar con el estado inicial; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "[1 2 3 4 5 6]") {
		t.Errorf("la animación debe terminar con el arreglo ordenado; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "Inserta") {
		t.Errorf("la animación debe rotular las inserciones; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "*") {
		t.Errorf("la animación debe marcar los elementos insertados; salida:\n%s", salida)
	}
}
