package main

// EsPalindromo indica si texto es un palíndromo.
//
// Se compara el primer carácter con el último; si son iguales se verifica
// el interior (texto[1:len(texto)-1]) recursivamente. Si son distintos
// devuelve false.
//
// Caso base: cadena vacía o de longitud 1 son palíndromos.
//
// Precondición: texto está formado por caracteres de un byte (ASCII).
func EsPalindromo(texto string) bool {
	if len(texto) <= 1 {
		return true
	}

	primero := texto[0]
	ultimo := texto[len(texto)-1]
	interior := texto[1 : len(texto)-1]

	return primero == ultimo && EsPalindromo(interior)
}
