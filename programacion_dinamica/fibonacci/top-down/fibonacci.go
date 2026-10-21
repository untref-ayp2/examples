package main

// Fibonacci devuelve el n-ésimo número de Fibonacci con memoización: igual
// que la versión recursiva, pero cada resultado se guarda en mem y no se
// recalcula.
//
// Precondición: n >= 0.
// Complejidad: O(n) en tiempo, O(n) en espacio.
func Fibonacci(n int) int {
	mem := make(map[int]int)
	// Los dos primeros valores son los casos base y ya se conocen.
	mem[0] = 0
	mem[1] = 1
	return fibonacci(n, mem)
}

// fibonacci es la recursión con memoria. Si n ya está calculado lo devuelve
// del mapa; si no, lo calcula como la suma de los dos anteriores y lo guarda.
//
// Termina porque cada llamada baja n en 1 o 2 y los casos base 0 y 1 están
// cargados en mem desde el principio.
func fibonacci(n int, mem map[int]int) int {
	key := n
	result, ok := mem[key]
	if ok {
		return result
	}
	mem[key] = fibonacci(n-1, mem) + fibonacci(n-2, mem)
	return mem[key]
}
