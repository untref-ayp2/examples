package main

import "testing"

func TestArrayVacio(t *testing.T) {
	array := []int{}
	ordenado := []int{}
	if !iguales(Mergesort(array), ordenado) {
		t.Errorf("Mergesort(%v) = %v; se esperaba %v", array, Mergesort(array), ordenado)
	}
}

func TestArrayUnElemento(t *testing.T) {
	array := []int{3}
	ordenado := []int{3}
	if !iguales(Mergesort(array), ordenado) {
		t.Errorf("Mergesort(%v) = %v; se esperaba %v", array, Mergesort(array), ordenado)
	}
}

func TestCantidadPar(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	ordenado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !iguales(Mergesort(array), ordenado) {
		t.Errorf("Mergesort(%v) = %v; se esperaba %v", array, Mergesort(array), ordenado)
	}
}

func TestCantidadImpar(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 6}
	ordenado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !iguales(Mergesort(array), ordenado) {
		t.Errorf("Mergesort(%v) = %v; se esperaba %v", array, Mergesort(array), ordenado)
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
