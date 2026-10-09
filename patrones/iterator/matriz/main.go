package main

import "fmt"

func main() {
	// Crea una matriz de ejemplo.
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matriz := NewMatriz(data)

	// Recorre la matriz con el iterador.
	it := matriz.Iterador()
	for it.Siguiente() {
		fmt.Println(it.Valor())
	}
	// Salida:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}
