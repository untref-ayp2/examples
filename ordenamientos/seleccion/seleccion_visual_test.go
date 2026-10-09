package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestAnimarSeleccion(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	original := append([]int(nil), arr...)

	var buf bytes.Buffer
	animarSeleccion(&buf, arr, 0)

	salida := buf.String()
	if !iguales(arr, original) {
		t.Errorf("la animación no debe modificar el arreglo original: quedó %v", arr)
	}
	if !strings.Contains(salida, "Estado inicial") {
		t.Errorf("la animación debe arrancar con el estado inicial; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "[11 12 22 25 64]") {
		t.Errorf("la animación debe terminar con el arreglo ordenado; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "coloca el mínimo") {
		t.Errorf("la animación debe rotular cada mínimo colocado; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "*") {
		t.Errorf("la animación debe marcar el mínimo y su posición original; salida:\n%s", salida)
	}
}
