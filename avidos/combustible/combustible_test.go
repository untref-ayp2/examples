package main

import (
	"testing"
)

func TestDiapositiva(t *testing.T) {
	distancia := 950
	capacidad := 400
	estaciones := []int{200, 375, 550, 750}

	got := MinimasRecargas(distancia, capacidad, estaciones)
	if got != 2 {
		t.Errorf("MinimasRecargas(%d, %d, %v) = %d; se esperaba 2", distancia, capacidad, estaciones, got)
	}
}

func TestLlegaConLaCargaInicial(t *testing.T) {
	distancia := 950
	capacidad := 950
	estaciones := []int{200, 375, 550, 750}

	got := MinimasRecargas(distancia, capacidad, estaciones)
	if got != 0 {
		t.Errorf("MinimasRecargas(%d, %d, %v) = %d; se esperaba 0", distancia, capacidad, estaciones, got)
	}
}

func TestLlegaParandoEnTodas(t *testing.T) {
	distancia := 950
	capacidad := 200
	estaciones := []int{200, 375, 550, 750}

	got := MinimasRecargas(distancia, capacidad, estaciones)
	if got != 4 {
		t.Errorf("MinimasRecargas(%d, %d, %v) = %d; se esperaba 4", distancia, capacidad, estaciones, got)
	}
}

func TestLlegaParandoSolamenteEnLaUltima(t *testing.T) {
	distancia := 950
	capacidad := 800
	estaciones := []int{200, 375, 550, 750, 950}

	got := MinimasRecargas(distancia, capacidad, estaciones)
	if got != 1 {
		t.Errorf("MinimasRecargas(%d, %d, %v) = %d; se esperaba 1", distancia, capacidad, estaciones, got)
	}
}

// TestParadasDeRecarga verifica que la estrategia ávida elige la estación más
// lejana alcanzable en cada tramo.
func TestParadasDeRecarga(t *testing.T) {
	distancia := 950
	capacidad := 400
	estaciones := []int{200, 375, 550, 750}

	got := paradasDeRecarga(distancia, capacidad, estaciones)
	esperado := []int{375, 750}
	if !mismasParadas(got, esperado) {
		t.Errorf("paradasDeRecarga(%d, %d, %v) = %v; se esperaba %v", distancia, capacidad, estaciones, got, esperado)
	}
}

// mismasParadas indica si dos listas de paradas son iguales, en el mismo
// orden.
func mismasParadas(a, b []int) bool {
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
