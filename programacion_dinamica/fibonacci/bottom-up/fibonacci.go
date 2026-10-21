package main

// Fibonacci devuelve el n-ésimo número de Fibonacci con tabulación: arma la
// tabla table de abajo hacia arriba, empezando por los casos base y usando
// las dos celdas anteriores para completar cada una.
//
// Precondición: n >= 0.
// Complejidad: O(n) en tiempo, O(n) en espacio.
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	table := make([]int, n+1)
	table[1] = 1
	// table[0] ya es 0: la tabla arranca con los dos casos base.
	for i := 2; i < len(table); i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n]
}
