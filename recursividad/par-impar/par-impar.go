package main

// esPar decide si n es par por recursión mutua: 0 es par y, salvo ese caso,
// la paridad de n es la imparidad de n-1.
func esPar(n int) bool {
	if n == 0 {
		return true
	}
	return esImpar(n - 1)
}

// esImpar decide si n es impar por recursión mutua con esPar: 0 no es impar y,
// salvo ese caso, la imparidad de n es la paridad de n-1.
func esImpar(n int) bool {
	if n == 0 {
		return false
	}
	return esPar(n - 1)
}
