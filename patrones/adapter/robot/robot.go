package main

// Este archivo muestra Adapter con un robot que mide en metros y centímetros
// y un cliente que espera una interfaz Medible en pulgadas. RobotAdapter
// envuelve al Robot y hace la conversión: el cliente sigue trabajando con
// Medible y el Robot no se modifica.

// Medible es la interfaz que espera el cliente.
type Medible interface {
	Medir() float64
}

// Robot es el Adaptado: mide en metros y centímetros.
type Robot struct {
	metros      int
	centimetros int
}

// Medir devuelve la medida del robot en metros y centímetros.
func (r *Robot) Medir() (int, int) {
	return r.metros, r.centimetros
}

// RobotAdapter envuelve un Robot y adapta su interfaz a Medible.
type RobotAdapter struct {
	adaptado *Robot
}

// Medir adapta la medida del robot a la interfaz Medible: convierte metros y
// centímetros a pulgadas (1 pulgada = 2.54 cm).
func (a *RobotAdapter) Medir() float64 {
	metros, centimetros := a.adaptado.Medir()
	totalCentimetros := (metros * 100) + centimetros
	pulgadas := float64(totalCentimetros) / 2.54
	return pulgadas
}
