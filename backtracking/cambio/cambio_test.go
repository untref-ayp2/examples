package main

import "testing"

func TestCambioDeMoneda(t *testing.T) {
	casos := []struct {
		nombre         string
		cantidad       int
		denominaciones []int
		piezas         int // cantidad esperada de piezas de la mejor solución
	}{
		{"caso del ejemplo", 27, []int{1, 2, 9, 10, 20, 50}, 3},
		{"con moneda de 1", 7, []int{1, 5}, 3},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			solucion := CambioDeMoneda(caso.cantidad, caso.denominaciones)
			if len(solucion) != caso.piezas {
				t.Fatalf("CambioDeMoneda(%d, %v) = %v, se esperaban %d piezas",
					caso.cantidad, caso.denominaciones, solucion, caso.piezas)
			}
			if !sumaCorrecta(solucion, caso.cantidad) {
				t.Errorf("la suma de %v no da %d", solucion, caso.cantidad)
			}
		})
	}
}

func TestCambioDeMonedaSinSolucion(t *testing.T) {
	if sol := CambioDeMoneda(3, []int{5, 10}); sol != nil {
		t.Errorf("CambioDeMoneda(3, [5 10]) = %v, se esperaba nil", sol)
	}
}

func sumaCorrecta(solucion []int, cantidad int) bool {
	total := 0
	for _, moneda := range solucion {
		total += moneda
	}
	return total == cantidad
}
