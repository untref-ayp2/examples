package main

// Factorial calcula el factorial de n (n! = n × (n-1) × ... × 1) de forma
// recursiva.
//
// Caso base: 0! = 1 (y también 1! = 1, que se llega por la misma recursión).
// Para n > 0, Factorial(n) = n * Factorial(n-1). La llamada recursiva
// disminuye n en 1 en cada paso, por lo que el problema se achica y se llega
// al caso base.
//
// Precondición: n >= 0.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}
