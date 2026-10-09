package main

import "testing"

func TestOrdenarSeleccion(t *testing.T) {
	casos := []struct {
		nombre   string
		entrada  []int
		esperado []int
	}{
		{"vacío", []int{}, []int{}},
		{"un elemento", []int{3}, []int{3}},
		{"desordenado", []int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
		{"con repetidos", []int{3, 1, 3, 2, 1}, []int{1, 1, 2, 3, 3}},
	}
	for _, c := range casos {
		OrdenarSeleccion(c.entrada)
		if !iguales(c.entrada, c.esperado) {
			t.Errorf("%s: quedó %v; se esperaba %v", c.nombre, c.entrada, c.esperado)
		}
	}
}

// iguales compara dos slices de enteros elemento por elemento.
func iguales(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
