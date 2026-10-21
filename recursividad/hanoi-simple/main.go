package main

import "fmt"

func main() {
	// Para experimentar: con 3 o 4 discos se sigue cómodamente en pantalla;
	// con 10 son 1023 movimientos.
	n := 3

	movimientos := Mover(n, "A", "B", "C")

	fmt.Printf("Torres de Hanoi con %d discos: %d movimientos\n\n", n, len(movimientos))
	for i, m := range movimientos {
		fmt.Printf("%3d. disco %d: %s -> %s\n", i+1, m.Disco, m.Origen, m.Destino)
	}

	fmt.Printf("\n¡Listo! Todos los discos quedaron en la vara C.\n")
}
