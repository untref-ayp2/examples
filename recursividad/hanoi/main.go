package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// varaIndice traduce el nombre de una vara ("A", "B" o "C") al índice de la
// lista de varas que maneja la capa visual.
func varaIndice(nombre string) int {
	switch nombre {
	case "A":
		return 0
	case "B":
		return 1
	default:
		return 2
	}
}

func main() {
	const delay = 400 * time.Millisecond

	estado := nuevoEstadoHanoi(LeerCantidadDiscos(bufio.NewReader(os.Stdin), os.Stdout))

	estado.dibujar(os.Stdout, 0)

	Mover(estado.discos, "A", "B", "C", func(disco int, origen, destino string) {
		estado.mover(disco, varaIndice(origen), varaIndice(destino))
		estado.dibujar(os.Stdout, delay)
	})

	fmt.Println("¡Listo! Todos los discos quedaron en la vara C.")
}
