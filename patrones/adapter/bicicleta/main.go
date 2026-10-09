package main

import "fmt"

func main() {
	// Instanciar bicicleta playera y de carrera
	playera := &Playera{}
	carreras := &BicicletaDeCarreras{}

	// Frenar la bicicleta playera y de carrera utilizando sus propios frenos
	playera.PedalearHaciaAtras()
	carreras.Frenar()

	fmt.Println("--------------------")

	bicicletas := []Bicicleta{&BicicletaDeCarreras{}, &PlayeraAdapter{}}

	bicicletas[0].Frenar()
	bicicletas[1].Frenar()

	fmt.Println("--------------------")

	// Utilizar el adaptador de freno para frenar la bicicleta playera como si fuera una bicicleta de carrera
	adapter := &PlayeraAdapter{}
	adapter.Frenar()
}
