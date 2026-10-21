package main

// Rob devuelve la máxima cantidad de dinero que se puede robar en una calle
// de casas, sin robar dos casas vecinas.
//
// Misma decisión que la versión recursiva (robar la casa actual y saltear la
// siguiente, o no robar y pasar a la siguiente), pero con memoización de los
// estados ya resueltos.
//
// Precondición: houses puede estar vacío (devuelve 0).
// Complejidad: O(n) en tiempo, O(n) en espacio.
func Rob(houses []int) int {
	return rob(houses, 0, make(map[int]int))
}

// rob es la recursión con memoria: la clave es la posición desde donde se
// sigue decidiendo. Caso base: no quedan casas.
func rob(houses []int, current int, mem map[int]int) int {
	key := current
	if result, ok := mem[key]; ok {
		return result
	}

	if current >= len(houses) {
		mem[key] = 0
		return mem[key]
	}

	mem[key] = max(houses[current]+rob(houses, current+2, mem), rob(houses, current+1, mem))
	return mem[key]
}
