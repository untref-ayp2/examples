package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// para experimentar: cambiá la distancia, la capacidad del tanque o las
	// estaciones de servicio (en orden de aparición en la ruta)
	distancia := 950
	capacidad := 400
	estaciones := []int{200, 375, 550, 750}
	const delay = 300 * time.Millisecond

	paradas := paradasDeRecarga(distancia, capacidad, estaciones)
	recargas := MinimasRecargas(distancia, capacidad, estaciones)

	fmt.Printf("Recorrido de %d km con tanque de %d km de autonomía\n", distancia, capacidad)
	fmt.Printf("Estaciones de servicio: %v\n\n", estaciones)
	dibujarRecorrido(os.Stdout, distancia, estaciones, paradas, delay)

	if len(paradas) > 0 {
		fmt.Printf("\nRecargas mínimas: %d (paradas en %v)\n", recargas, paradas)
	} else {
		fmt.Printf("\nRecargas mínimas: %d (no hace falta recargar)\n", recargas)
	}
}
