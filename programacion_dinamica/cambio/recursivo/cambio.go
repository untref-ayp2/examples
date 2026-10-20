package main

import "math"

// Cambio devuelve la cantidad mínima de monedas (tomadas de las
// denominaciones en coins) necesarias para armar el monto exacto.
// Si el monto no se puede armar con esas monedas, devuelve -1.
func Cambio(monto int, coins []int) int {
	if monto == 0 {
		return 0
	}

	menor := int(math.MaxInt32) //simula infinito
	for _, coin := range coins {
		if monto >= coin {
			// un resultado negativo significa que por esa moneda no hay
			// solución; no lo tengo en cuenta
			resto := Cambio(monto-coin, coins)
			if resto >= 0 {
				menor = min(menor, resto)
			}
		}
	}

	// si ninguna moneda lleva a una solución, el monto no se puede armar
	if menor == int(math.MaxInt32) {
		return -1
	}
	return 1 + menor
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
