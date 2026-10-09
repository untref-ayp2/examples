package main

import "testing"

func iguales(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}

func TestCambio(t *testing.T) {
	cantidad := 1234
	esperado := map[int]int{
		1000: 1,
		200:  1,
		20:   1,
		10:   1,
		2:    2,
	}

	got := Cambiar(cantidad)
	if !iguales(got, esperado) {
		t.Errorf("Cambiar(%d) = %v; se esperaba %v", cantidad, got, esperado)
	}
}

func TestDiapositivas(t *testing.T) {
	cantidad := 5528
	esperado := map[int]int{
		2000: 2,
		1000: 1,
		500:  1,
		20:   1,
		5:    1,
		2:    1,
		1:    1,
	}

	got := Cambiar(cantidad)
	if !iguales(got, esperado) {
		t.Errorf("Cambiar(%d) = %v; se esperaba %v", cantidad, got, esperado)
	}
}
