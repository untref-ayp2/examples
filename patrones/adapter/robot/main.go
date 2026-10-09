package main

import "fmt"

func main() {
	robot := &Robot{metros: 10, centimetros: 50}
	adaptado := &RobotAdapter{adaptado: robot}
	distancia := adaptado.Medir()
	fmt.Printf("Distancia: %.2f pulgadas\n", distancia)
	// Salida: Distancia: 413.39 pulgadas
}
