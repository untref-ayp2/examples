package main

import "math"

// Cambio devuelve la cantidad mínima de monedas (tomadas de las
// denominaciones en coins) necesarias para armar el monto exacto.
// Si el monto no se puede armar con esas monedas, devuelve -1.
//
// Versión bottom-up: en vez de resolver el problema «de arriba hacia abajo»
// con recursión, se arma una tabla con las respuestas de todos los montos
// desde 0 hasta monto. table[x] es la cantidad mínima de monedas para armar
// x, y se calcula a partir de los montos menores ya resueltos.
//
// Precondición: monto >= 0 y todas las denominaciones son positivas.
// Complejidad: O(monto * len(coins)) en tiempo, O(monto) en espacio.
func Cambio(monto int, coins []int) int {
	n := monto

	table := make([]int, n+1)
	table[0] = 0

	for x := 1; x <= n; x++ {
		menor := int(math.MaxInt32)
		for _, coin := range coins {
			if x >= coin {
				// Se mira la celda del subproblema que queda al sacar esta
				// moneda. Si esa celda está marcada con el centinela, la
				// comparación deja a menor en el centinela.
				menor = min(menor, table[x-coin])
			}
		}
		// Si ninguna moneda lleva a una celda resoluble, el monto x no se
		// puede armar: la celda queda marcada con el infinito.
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
