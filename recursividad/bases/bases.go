package main

import "fmt"

// ImprimirEnDecimal convierte n a su representación decimal en texto,
// imprimiendo los dígitos de izquierda a derecha.
//
// Es la versión con base fija 10 de ImprimirEnBase: primero se resuelve el
// número sin su último dígito (la parte recursiva) y después se agrega ese
// dígito al final.
//
// Caso base: n < 10 ya es un único dígito, así que no hace falta dividir.
// Termina porque cada llamada divide n por 10 y, en los enteros, eso lo
// achica de verdad (tarde o temprano llega a un valor < 10).
//
// Precondición: n >= 0.
func ImprimirEnDecimal(n int) string {
	res := ""

	if n >= 10 {
		res += ImprimirEnDecimal(n / 10)
	}

	return res + fmt.Sprintf("%d", n%10)
}

// digits son los 16 símbolos que se usan como dígitos. El alfabeto
// hexadecimal va del 0 al 9 y sigue con a..f, así que cada posición de la
// cadena es el dígito de ese valor: digits[10] == 'a', digits[15] == 'f'.
// Por eso la base máxima que soporta ImprimirEnBase es 16.
const digits string = "0123456789abcdef"

// ImprimirEnBase convierte n a su representación en la base dada, en texto.
//
// Se resuelve primero el número dividido por la base (las cifras más
// significativas) y después se agrega el resto como última cifra.
//
// Caso base: n < base deja una sola cifra (n % base == n). Termina porque
// cada llamada divide n por base >= 2, y la división entera achica n.
//
// Precondición: n >= 0 y 2 <= base <= 16 (fuera de ese rango no hay símbolo
// para la cifra en digits).
func ImprimirEnBase(n int, base int) string {
	res := ""

	if n >= base {
		res += ImprimirEnBase(n/base, base)
	}

	return res + string(digits[n%base])
}
