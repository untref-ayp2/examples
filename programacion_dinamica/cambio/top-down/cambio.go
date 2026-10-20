package main

import "math"

// Cambio devuelve la cantidad mínima de monedas (tomadas de las
// denominaciones en coins) necesarias para armar el monto exacto.
// Si el monto no se puede armar con esas monedas, devuelve -1.
func Cambio(monto int, coins []int) int {
	return cambio(monto, coins, make(map[int]int))
}

func cambio(monto int, coins []int, mem map[int]int) int {
	if monto == 0 {
		return 0
	}

	if result, ok := mem[monto]; ok {
		return result
	}

	menor := int(math.MaxInt32)
	for _, coin := range coins {
		if monto >= coin {
			// un resultado negativo significa que por esa moneda no hay
			// solución; no lo tengo en cuenta
			resto := cambio(monto-coin, coins, mem)
			if resto >= 0 {
				menor = min(menor, resto)
			}
		}
	}

	if menor == int(math.MaxInt32) {
		// ninguna moneda lleva a una solución: el monto es imposible
		mem[monto] = -1
	} else {
		mem[monto] = 1 + menor
	}
	return mem[monto]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
