package main

// SumaArray devuelve la suma de todos los elementos de arreglo, por división
// y conquista: parte el arreglo en dos mitades, suma cada una por separado y
// suma los dos resultados.
//
// Precondición: arreglo puede estar vacío (en ese caso la suma es 0).
// Complejidad: O(n).
func SumaArray(arreglo []int) int {
	// Casos base: sin elementos la suma es 0; con uno solo, ese elemento.
	if len(arreglo) == 0 {
		return 0
	}
	if len(arreglo) == 1 {
		return arreglo[0]
	}

	// Dividir: el punto medio. arreglo[:medio] y arreglo[medio:] son
	// «vistas» del mismo arreglo, no copias: por eso dividir no cuesta
	// memoria extra.
	medio := len(arreglo) / 2

	// Conquistar y combinar: la suma de las mitades es la suma del total.
	return SumaArray(arreglo[:medio]) + SumaArray(arreglo[medio:])
}
