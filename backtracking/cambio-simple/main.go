package main

import "fmt"

func main() {
	// Cantidad a cambiar y monedas disponibles. Para experimentar, probá
	// otras combinaciones: por ejemplo 30 con {1, 7, 10} o 3 con {5, 10}
	// (este último no tiene solución).
	cantidad := 27
	denominaciones := []int{1, 2, 9, 10, 20, 50}

	cambio := CambioDeMoneda(cantidad, denominaciones)

	if cambio == nil {
		fmt.Printf("No hay forma de cambiar %d con %v.\n", cantidad, denominaciones)
		return
	}

	fmt.Printf("Cambio de %d con %v:\n", cantidad, denominaciones)
	fmt.Printf("  %v (%d piezas)\n", cambio, len(cambio))
}
