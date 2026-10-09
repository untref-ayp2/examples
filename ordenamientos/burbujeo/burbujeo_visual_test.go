package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestAnimarBurbujeo(t *testing.T) {
	arr := []int{5, 1, 4, 2, 8}
	original := append([]int(nil), arr...)

	var buf bytes.Buffer
	animarBurbujeo(&buf, arr, 0)

	salida := buf.String()
	if !iguales(arr, original) {
		t.Errorf("la animación no debe modificar el arreglo original: quedó %v", arr)
	}
	if !strings.Contains(salida, "Estado inicial") {
		t.Errorf("la animación debe arrancar con el estado inicial; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "[1 2 4 5 8]") {
		t.Errorf("la animación debe terminar con el arreglo ordenado; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "intercambia") {
		t.Errorf("la animación debe rotular los intercambios; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "*") {
		t.Errorf("la animación debe marcar los elementos intercambiados; salida:\n%s", salida)
	}
}
