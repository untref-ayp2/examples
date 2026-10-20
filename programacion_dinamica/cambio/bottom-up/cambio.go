package main

import "math"

// Cambio devuelve la cantidad mínima de monedas (tomadas de las
// denominaciones en coins) necesarias para armar el monto exacto.
// Si el monto no se puede armar con esas monedas, devuelve -1.
func Cambio(monto int, coins []int) int {
	n := monto

	table := make([]int, n+1)
	table[0] = 0

	for x := 1; x <= n; x++ {
		menor := int(math.MaxInt32)
		for _, coin := range coins {
			if x >= coin {
				menor = min(menor, table[x-coin])
			}
		}
		// si ninguna moneda lleva a una celda resoluble, el monto x no se
		// puede armar: la celda queda marcada con el infinito
		if menor == int(math.MaxInt32) {
			table[x] = int(math.MaxInt32)
		} else {
			table[x] = 1 + menor
		}
	}
	if table[n] == int(math.MaxInt32) {
		return -1
	}
	return table[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
