package main

import "fmt"

type Bicicleta interface {
	Frenar()
}

// LEGACY
// Implementación de la bicicleta playera
type Playera struct{}

// LEGACY
// Frenar la bicicleta playera
func (p *Playera) PedalearHaciaAtras() {
	fmt.Println("Frenando la bicicleta playera...")
}

// Implementación de la bicicleta de carrera
type BicicletaDeCarreras struct{}

// Frenar la bicicleta de carrera
func (c *BicicletaDeCarreras) Frenar() {
	fmt.Println("Frenando la bicicleta de carrera...")
}

// Adaptador de freno que permite utilizar el freno de la bicicleta playera como si fuera una bicicleta de carrera
type PlayeraAdapter struct {
	Playera
}

// Frenar la bicicleta playera como si fuera una bicicleta de carrera
func (b *PlayeraAdapter) Frenar() {
	fmt.Println("Adaptando el freno de la bicicleta playera para la bicicleta de carrera...")
	fmt.Print(" -> ")
	b.PedalearHaciaAtras()
}
