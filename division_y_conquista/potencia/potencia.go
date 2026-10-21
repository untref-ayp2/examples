package main

// Potencia calcula n elevado a la x por división y conquista: en vez de
// multiplicar n por sí mismo x veces, parte el exponente a la mitad y
// reutiliza esa potencia.
//
// Precondición: x >= 0 (con exponente negativo el resultado no sería entero).
// Complejidad: O(log x) multiplicaciones.
func Potencia(n int, x int) int {
	// Caso base 1: todo número elevado a 0 es 1.
	if x == 0 {
		return 1
	}

	// Caso base 2: elevado a 1 es el mismo número.
	if x == 1 {
		return n
	}

	if x%2 == 0 {
		// Exponente par: n^x = (n^(x/2))^2. La media potencia se calcula
		// una sola vez y se multiplica por sí misma; si se escribiera
		// Potencia(n, x/2) * Potencia(n, x/2) se calcularía dos veces.
		mediaPotencia := Potencia(n, x/2)
		return mediaPotencia * mediaPotencia
	}

	// Exponente impar: n^x = n * (n^((x-1)/2))^2, porque (x-1)/2 es entero.
	mediaPotencia := Potencia(n, (x-1)/2)
	return n * mediaPotencia * mediaPotencia
}
