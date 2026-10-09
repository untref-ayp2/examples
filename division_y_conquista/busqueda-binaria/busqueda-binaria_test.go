package main

import "testing"

func TestBuscandoElementoExistente(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 8, 9, 12, 27, 50, 70}
	pos := BusquedaBinaria(array, 27)
	if pos != 8 {
		t.Errorf("BusquedaBinaria(%v, 27) = %d; se esperaba 8", array, pos)
	}
}

func TestBuscandoElementoInexistente(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 8, 9, 12, 27, 50, 70}
	pos := BusquedaBinaria(array, 23)
	if pos != -1 {
		t.Errorf("BusquedaBinaria(%v, 23) = %d; se esperaba -1", array, pos)
	}
}
