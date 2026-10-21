package main

// Rob devuelve la máxima cantidad de dinero que se puede robar en una calle
// de casas, sin robar dos casas vecinas.
//
// La recurrencia es una decisión en cada casa i:
//
//	Rob(i) = max(houses[i] + Rob(i+2),  // robar acá y saltear la siguiente
//	             Rob(i+1))              // no robar acá
//
// Caso base: i >= len(houses) (no quedan casas) devuelve 0. Termina porque
// cada llamada avanza a una casa posterior.
//
// Precondición: houses puede estar vacío (devuelve 0).
// Complejidad: exponencial, acotada por O(2^n): cada casa ramifica en dos y
// se recalculan ramas.
func Rob(houses []int) int {
	return rob(houses, 0)
}

// rob es la recursión: evalúa la decisión de robar o no la casa current.
func rob(houses []int, current int) int {
	if current >= len(houses) {
		return 0
	}

	return max(houses[current]+rob(houses, current+2), rob(houses, current+1))
}
