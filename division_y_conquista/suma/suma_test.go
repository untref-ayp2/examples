package main

import "testing"

func TestSumaElementos(t *testing.T) {
	arreglo := []int{1, 2, 3, 4, 5, 6}
	got := SumaArray(arreglo)
	if got != 21 {
		t.Errorf("SumaArray(%v) = %d; se esperaba 21", arreglo, got)
	}
}

func TestSumaVacio(t *testing.T) {
	arreglo := []int{}
	got := SumaArray(arreglo)
	if got != 0 {
		t.Errorf("SumaArray(%v) = %d; se esperaba 0", arreglo, got)
	}
}

func TestSumaOpuestos(t *testing.T) {
	arreglo := []int{-2, -1, 0, 1, 2}
	got := SumaArray(arreglo)
	if got != 0 {
		t.Errorf("SumaArray(%v) = %d; se esperaba 0", arreglo, got)
	}
}
