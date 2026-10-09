package main

import "testing"

func TestArrayVacioQ(t *testing.T) {
	array := []int{}
	ordenado := []int{}
	if !iguales(Quicksort(array), ordenado) {
		t.Errorf("Quicksort(%v) = %v; se esperaba %v", array, Quicksort(array), ordenado)
	}
}

func TestArrayUnElementoQ(t *testing.T) {
	array := []int{3}
	ordenado := []int{3}
	if !iguales(Quicksort(array), ordenado) {
		t.Errorf("Quicksort(%v) = %v; se esperaba %v", array, Quicksort(array), ordenado)
	}
}

func TestCantidadParQ(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	ordenado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !iguales(Quicksort(array), ordenado) {
		t.Errorf("Quicksort(%v) = %v; se esperaba %v", array, Quicksort(array), ordenado)
	}
}

func TestCantidadImparQ(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 6}
	ordenado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !iguales(Quicksort(array), ordenado) {
		t.Errorf("Quicksort(%v) = %v; se esperaba %v", array, Quicksort(array), ordenado)
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
