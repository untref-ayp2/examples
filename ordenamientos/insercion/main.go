package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// para experimentar: cambiá el arreglo o el delay (ponelo en 0 para ver
	// la animación completa de un solo tirón, sin pausas)
	arr := []int{5, 2, 4, 6, 1, 3}
	const delay = 300 * time.Millisecond

	fmt.Print("Inserción — estado del arreglo paso a paso\n\n")

	fmt.Printf("Antes:  %v\n\n", arr)
	animarInsercion(os.Stdout, arr, delay)

	OrdenarInsercion(arr)
	fmt.Printf("Después: %v\n", arr)
}
