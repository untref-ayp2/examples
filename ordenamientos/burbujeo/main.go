package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// para experimentar: cambiá el arreglo o el delay (ponelo en 0 para ver
	// la animación completa de un solo tirón, sin pausas)
	arr := []int{5, 1, 4, 2, 8}
	const delay = 300 * time.Millisecond

	fmt.Print("Burbujeo — estado del arreglo paso a paso\n\n")

	fmt.Printf("Antes:  %v\n\n", arr)
	animarBurbujeo(os.Stdout, arr, delay)

	OrdenarBurbujeo(arr)
	fmt.Printf("Después: %v\n", arr)
}
