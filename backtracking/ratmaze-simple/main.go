package main

import "fmt"

func main() {
	// Laberinto de prueba: true es celda libre y false es pared. Para
	// experimentar, cambiá el mapa o mirá cómo cambia el camino si se
	// reordenan las direcciones dentro de backtracking.
	maze := [][]bool{
		{true, true, true, true, true, true, true, true, true, true, true, true, true},
		{false, false, true, false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false, false, false, false},
		{false, false, true, true, true, true, true, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false, false, false, false},
		{false, false, true, true, true, true, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false, false, false, false},
		{true, true, true, true, true, true, true, true, true, true, true, true, true},
	}

	ratmaze := NuevoRatMaze(maze)

	if ratmaze.Resolver() {
		fmt.Println("Laberinto resuelto (el camino se ve con 'o'):")
	} else {
		fmt.Println("No existe un camino hasta el destino:")
	}
	fmt.Print(ratmaze.Tablero())
}
