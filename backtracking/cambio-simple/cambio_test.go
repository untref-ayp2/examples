package main

import "testing"

// TestCambioDeMoneda comprueba que la solución tenga la menor cantidad de
// piezas posible y que las monedas sumen exactamente la cantidad pedida.
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

// TestCambioDeMonedaSinSolucion comprueba que, cuando ninguna combinación
// suma exactamente la cantidad, la función devuelva nil.
func TestCambioDeMonedaSinSolucion(t *testing.T) {
	if sol := CambioDeMoneda(3, []int{5, 10}); sol != nil {
		t.Errorf("CambioDeMoneda(3, [5 10]) = %v, se esperaba nil", sol)
	}
}

// TestCambioDeMonedaCero comprueba el borde de la cantidad cero: el cambio
// vacío es la única solución (no hacen falta monedas).
func TestCambioDeMonedaCero(t *testing.T) {
	if sol := CambioDeMoneda(0, []int{1, 5}); len(sol) != 0 {
		t.Errorf("CambioDeMoneda(0, [1 5]) = %v, se esperaba un cambio vacío", sol)
	}
}

// sumaCorrecta indica si las monedas de la solución suman exactamente total.
func sumaCorrecta(solucion []int, total int) bool {
	suma := 0
	for _, moneda := range solucion {
		suma += moneda
	}
	return suma == total
}
