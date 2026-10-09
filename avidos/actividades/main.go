package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// para experimentar: agregá o quitá actividades. La lista debe estar
	// ordenada por tiempo de finalización.
	actividades := []Actividad{
		{"a1", 1, 4},
		{"a2", 3, 5},
		{"a3", 0, 6},
		{"a4", 5, 7},
		{"a5", 3, 9},
		{"a6", 5, 9},
		{"a7", 6, 10},
		{"a8", 8, 11},
		{"a9", 8, 12},
		{"a10", 2, 14},
		{"a11", 12, 15},
	}
	const delay = 400 * time.Millisecond

	seleccionadas := SelectorActividadesIterativo(actividades)

	fmt.Printf("Actividades: %d. Seleccionadas: %d\n\n", len(actividades), len(seleccionadas))
	dibujarLineaTiempo(os.Stdout, seleccionadas, delay)
	fmt.Printf("\nElegidas: ")
	for i, a := range seleccionadas {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%s (%d-%d)", a.Nombre, a.Inicio, a.Fin)
	}
	fmt.Println()
}
