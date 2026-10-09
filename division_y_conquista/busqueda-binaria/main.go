package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 8, 9, 12, 27, 50, 70}

	// para experimentar: cambiá el array o los elementos a buscar
	busquedas := []int{27, 23, 1, 70, 100}
	for _, x := range busquedas {
		pos := BusquedaBinaria(array, x)
		fmt.Printf("BusquedaBinaria(%v, %d) = %d\n", array, x, pos)
	}
}
