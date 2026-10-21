package main

import "fmt"

func main() {
	// Para experimentar: probá distintos n. Con 2 y 3 no hay solución; con 4
	// hay 2; con 5 hay 10; con 8 hay 92. Hasta ahí NReinasTodas tarda menos de
	// un milisegundo: el crecimiento fuerte recién se nota con 10 (milésimas
	// de segundo) y con 12 (decenas de milisegundos), porque la cantidad de
	// soluciones crece muy rápido.
	n := 6

	soluciones := NReinasTodas(n)

	fmt.Printf("Tablero de %dx%d: %d soluciones\n\n", n, n, len(soluciones))
	if len(soluciones) == 0 {
		fmt.Println("No hay ninguna forma de ubicar las reinas.")
		return
	}

	// Se muestra la primera solución y, como referencia, cuántas hay.
	fmt.Printf("Primera solución (columna de la reina de cada fila):\n  %v\n\n", NReinas(n))
	fmt.Println("Tablero de la primera solución (# es reina, . es vacío):")
	dibujarTablero(n, soluciones[0])
}

// dibujarTablero imprime un tablero n x n con las reinas de solucion, donde
// solucion[i] es la columna de la reina de la fila i.
func dibujarTablero(n int, solucion []int) {
	for fila := 0; fila < n; fila++ {
		for columna := 0; columna < n; columna++ {
			if solucion[fila] == columna {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
