package main

// Rob devuelve la máxima cantidad de dinero que se puede robar en una calle
// de casas, sin robar dos casas vecinas.
//
// Versión bottom-up sin tabla: como cada celda depende solo de las dos
// anteriores, alcanza con ir guardando los dos últimos valores (en vez de
// toda la tabla). Mismo resultado que la versión con tabla, con O(1) de
// espacio: las dos aplican la misma decisión
// max(houses[x]+table[x-2], table[x-1]) en cada paso, y los casos base
// coinciden (table[0] = houses[0] y table[1] = max(houses[0], houses[1])).
// La única diferencia es que acá no se guarda toda la tabla, porque cada
// celda nueva solo necesita las dos anteriores.
//
// Precondición: houses puede estar vacío (devuelve 0).
// Complejidad: O(n) en tiempo, O(1) en espacio.
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

	// prevPrev y prev son table[x-2] y table[x-1]. Ojo con el arranque: prev
	// representa table[1], y table[1] = max(houses[0], houses[1]) (con una o
	// dos casas se elige la que rinde más, porque son vecinas). No alcanza con
	// houses[1]: si la primera casa rinde más que la segunda, esa elección se
	// pierde y el resultado puede dar de menos.
	prevPrev := houses[0]
	prev := max(houses[0], houses[1])
	cur := 0

	for x := 2; x < n; x++ {
		cur = max(houses[x]+prevPrev, prev)

		prevPrev = prev
		prev = cur
	}

	return cur
}
