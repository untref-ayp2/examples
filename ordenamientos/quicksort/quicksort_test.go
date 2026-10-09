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

func TestElementosRepetidosQ(t *testing.T) {
	casos := []struct {
		array    []int
		ordenado []int
	}{
		{[]int{2, 2, 2, 2}, []int{2, 2, 2, 2}},
		{[]int{5, 5, 5}, []int{5, 5, 5}},
		{[]int{3, 1, 3, 2, 3}, []int{1, 2, 3, 3, 3}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
	}
	for _, caso := range casos {
		got := Quicksort(caso.array)
		if !iguales(got, caso.ordenado) {
			t.Errorf("Quicksort(%v) = %v; se esperaba %v", caso.array, got, caso.ordenado)
		}
	}
}

func TestQuicksortInPlace(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	ordenado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	Quicksort(array)

	if !iguales(array, ordenado) {
		t.Errorf("el slice de entrada debe quedar ordenado in-place; quedó %v", array)
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
