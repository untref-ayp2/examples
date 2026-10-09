package main

// Medible es la interfaz que espera el cliente.
type Medible interface {
	Medir() float64
}

// Robot es el Adaptado: mide en metros y centímetros.
type Robot struct {
	metros      int
	centimetros int
}

func (r *Robot) Medir() (int, int) {
	return r.metros, r.centimetros
}

// RobotAdapter envuelve un Robot y adapta su interfaz a Medible.
type RobotAdapter struct {
	adaptado *Robot
}

func (a *RobotAdapter) Medir() float64 {
	metros, centimetros := a.adaptado.Medir()
	totalCentimetros := (metros * 100) + centimetros
	pulgadas := float64(totalCentimetros) / 2.54
	return pulgadas
}
