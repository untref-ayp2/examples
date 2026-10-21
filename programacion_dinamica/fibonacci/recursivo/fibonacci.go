package main

// Fibonacci devuelve el n-ésimo número de Fibonacci con la recurrencia
// directa:
//
//	F(0) = 0, F(1) = 1
//	F(n) = F(n-1) + F(n-2)
//
// Casos base: n < 2 devuelve n (0 o 1). Termina porque cada llamada reduce n.
//
// Precondición: n >= 0.
// Complejidad: exponencial, acotada por O(2^n): la misma rama se recalcula
// una y otra vez (por eso después se ve la versión con memoria).
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
