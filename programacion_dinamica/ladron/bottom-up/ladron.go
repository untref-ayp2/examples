package main

// Rob devuelve la máxima cantidad de dinero que se puede robar en una calle
// de casas, sin robar dos casas vecinas.
//
// Versión bottom-up: table[x] guarda lo máximo que se puede robar
// considerando las primeras x+1 casas. Cada celda se completa con la misma
// decisión de siempre, mirando las celdas ya calculadas.
//
// Precondición: houses puede estar vacío (devuelve 0).
// Complejidad: O(n) en tiempo, O(n) en espacio.
func Rob(houses []int) int {
	n := len(houses)

	// Casos base: sin casas no se roba nada; con una o dos, se elige la que
	// rinde más, porque no se pueden robar dos casas vecinas.
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return max(houses[0], houses[n-1])
	}

	table := make([]int, n)
	table[0] = houses[0]
	table[1] = max(houses[0], houses[1])

	for x := 2; x < n; x++ {
		// Robar la casa x (y entonces saltear la x-1) o no robarla.
		table[x] = max(houses[x]+table[x-2], table[x-1])
	}

	return table[n-1]
}
