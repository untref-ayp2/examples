package main

import "testing"

func TestEsPar(t *testing.T) {
	pares := []int{0, 2, 4, 10}
	for _, n := range pares {
		if !esPar(n) {
			t.Errorf("esPar(%d) = false; se esperaba true", n)
		}
	}

	impares := []int{1, 3, 7}
	for _, n := range impares {
		if esPar(n) {
			t.Errorf("esPar(%d) = true; se esperaba false", n)
		}
	}
}

func TestEsImpar(t *testing.T) {
	impares := []int{1, 3, 7}
	for _, n := range impares {
		if !esImpar(n) {
			t.Errorf("esImpar(%d) = false; se esperaba true", n)
		}
	}

	pares := []int{0, 2, 4}
	for _, n := range pares {
		if esImpar(n) {
			t.Errorf("esImpar(%d) = true; se esperaba false", n)
		}
	}
}
