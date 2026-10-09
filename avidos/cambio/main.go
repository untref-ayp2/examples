package main

import (
	"fmt"
	"sort"
)

func main() {
	// para experimentar: cambiá la cantidad
	cantidad := 1234

	cambio := Cambiar(cantidad)
	fmt.Printf("Cambio de %d:\n", cantidad)

	denominaciones := make([]int, 0, len(cambio))
	total := 0
	for d, cantidadDenominacion := range cambio {
		denominaciones = append(denominaciones, d)
		total += cantidadDenominacion
	}
	sort.Sort(sort.Reverse(sort.IntSlice(denominaciones)))

	for _, d := range denominaciones {
		fmt.Printf("  %d de %d\n", cambio[d], d)
	}
	fmt.Printf("Total de piezas: %d\n", total)
}
