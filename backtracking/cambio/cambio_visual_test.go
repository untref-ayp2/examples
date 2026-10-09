package main

import (
	"strings"
	"testing"
)

func TestCambioDeMonedaConPasoObservaEventos(t *testing.T) {
	intentos := 0
	soluciones := 0
	solucion := CambioDeMonedaConPaso(6, []int{1, 2, 5}, ObservadorCambio{
		Intento: func(resto int, parcial []int) {
			intentos++
		},
		Solucion: func(solucion, mejor []int) {
			soluciones++
		},
	})

	if intentos == 0 || soluciones == 0 {
		t.Fatalf("el observador recibió %d intentos y %d soluciones, se esperaban más de 0",
			intentos, soluciones)
	}
	if len(solucion) != 2 {
		t.Errorf("CambioDeMonedaConPaso(6, [1 2 5]) = %v, se esperaban 2 piezas", solucion)
	}
}

func TestAnimarCambioMuestraSolucion(t *testing.T) {
	var salida strings.Builder
	AnimarCambio(6, []int{1, 2, 5}, &salida, 0)

	texto := salida.String()
	if !strings.Contains(texto, "Cambio de 6") {
		t.Errorf("la animación debería mostrar el monto a cambiar:\n%s", texto)
	}
	if !strings.Contains(texto, "El cambio con menos piezas") {
		t.Errorf("la animación debería mostrar la solución final:\n%s", texto)
	}
}
