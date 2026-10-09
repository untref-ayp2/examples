package main

import "fmt"

func main() {
	// para experimentar: agregá o cambiá los slices
	arreglos := [][]int{
		{1, 2, 3, 4, 5, 6}, // varios elementos
		{},                 // vacío
		{7},                // un solo elemento
		{-2, -1, 0, 1, 2},  // opuestos
	}
	for _, a := range arreglos {
		fmt.Printf("SumaArray(%v) = %d\n", a, SumaArray(a))
	}
}
