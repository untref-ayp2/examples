package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// para experimentar: cambiá el arreglo o el delay (ponelo en 0 para ver
	// la animación de una sola pasada)
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	const delay = 300 * time.Millisecond

	fmt.Printf("Mergesort — estado del arreglo paso a paso\n\n")

	fmt.Printf("Antes:  %v\n\n", array)
	animarMergesort(os.Stdout, array, delay)

	ordenado := Mergesort(array)
	fmt.Printf("Después: %v\n", ordenado)
}
