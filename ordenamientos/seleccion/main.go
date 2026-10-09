package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// para experimentar: cambiá el arreglo o el delay (ponelo en 0 para ver
	// la animación de una sola pasada)
	arr := []int{64, 25, 12, 22, 11}
	const delay = 300 * time.Millisecond

	fmt.Printf("Selección — estado del arreglo paso a paso\n\n")

	fmt.Printf("Antes:  %v\n\n", arr)
	animarSeleccion(os.Stdout, arr, delay)

	OrdenarSeleccion(arr)
	fmt.Printf("Después: %v\n", arr)
}
