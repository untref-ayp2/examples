package main

import "math"

// Cambio devuelve la cantidad mínima de monedas (tomadas de las
// denominaciones en coins) necesarias para armar el monto exacto.
// Si el monto no se puede armar con esas monedas, devuelve -1.
//
// Solución recursiva directa, sin memoización: la recurrencia es
//
//	Cambio(0) = 0
//	Cambio(monto) = 1 + min(Cambio(monto-coin))  sobre las coins <= monto
//
// Es decir: para armar un monto se elige la primera moneda y se resuelve
// recursivamente lo que queda. La respuesta es esa moneda más la mejor
// solución del resto.
//
// Precondición: monto >= 0 y todas las denominaciones son positivas.
// Complejidad: exponencial (recalcula los mismos submontos muchas veces).
func Cambio(monto int, coins []int) int {
	// Caso base: no queda nada por cambiar, hacen falta 0 monedas.
	if monto == 0 {
		return 0
	}

	// Se arranca con «infinito» para que cualquier solución real lo mejore.
	menor := int(math.MaxInt32)
	for _, coin := range coins {
		if monto >= coin {
			// Se prueba armar el resto sin esta moneda. Un resultado
			// negativo significa que por esa rama no hay solución: no se
			// tiene en cuenta.
			resto := Cambio(monto-coin, coins)
			if resto >= 0 {
				menor = min(menor, resto)
			}
		}
	}

	// Si ninguna moneda lleva a una solución, el monto no se puede armar.
	if menor == int(math.MaxInt32) {
		return -1
	}
	return 1 + menor
}
