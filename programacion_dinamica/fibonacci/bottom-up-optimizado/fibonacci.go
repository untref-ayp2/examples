package main

// Fibonacci devuelve el n-ésimo número de Fibonacci sin guardar la tabla
// completa: como cada valor depende solo de los dos anteriores, alcanza con
// llevar dos variables en vez del slice entero.
//
// Precondición: n >= 0.
// Complejidad: O(n) en tiempo, O(1) en espacio.
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	// prev2 y prev son F(i-2) y F(i-1); empiezan siendo F(0) y F(1).
	prev2 := 0
	prev := 1
	var current int

	for i := 2; i <= n; i++ {
		current = prev2 + prev
		prev2 = prev
		prev = current
	}
	return current
}
