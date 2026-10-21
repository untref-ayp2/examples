package main

import "math"

// Cambio devuelve la cantidad mínima de monedas (tomadas de las
// denominaciones en coins) necesarias para armar el monto exacto.
// Si el monto no se puede armar con esas monedas, devuelve -1.
//
// Misma recurrencia que la versión recursiva, pero con memoización: cada
// submonte se calcula una sola vez y se guarda en mem, así los cálculos
// repetidos se responden al instante.
//
// Precondición: monto >= 0 y todas las denominaciones son positivas.
// Complejidad: O(monto * len(coins)) en tiempo, O(monto) en espacio.
func Cambio(monto int, coins []int) int {
	return cambio(monto, coins, make(map[int]int))
}

// cambio es la recursión con memoria. Caso base: monto == 0 no necesita
// monedas. Termina porque cada llamada reduce el monto (las monedas son
// positivas) y los submontos repetidos salen de mem.
func cambio(monto int, coins []int, mem map[int]int) int {
	// Caso base: no queda nada por cambiar, hacen falta 0 monedas.
	if monto == 0 {
		return 0
	}

	// Si este submonte ya se resolvió, se devuelve lo guardado.
	if result, ok := mem[monto]; ok {
		return result
	}

	menor := int(math.MaxInt32)
	for _, coin := range coins {
		if monto >= coin {
			// Un resultado negativo significa que por esa moneda no hay
			// solución; no se tiene en cuenta.
			resto := cambio(monto-coin, coins, mem)
			if resto >= 0 {
				menor = min(menor, resto)
			}
		}
	}

	if menor == int(math.MaxInt32) {
		// Ninguna moneda lleva a una solución: el monto es imposible.
		mem[monto] = -1
	} else {
		mem[monto] = 1 + menor
	}
	return mem[monto]
}
