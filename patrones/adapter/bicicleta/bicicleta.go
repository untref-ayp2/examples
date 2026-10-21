package main

// Este archivo muestra el patrón Adapter con dos tipos de bicicleta que
// frenan de maneras distintas. El cliente quiere frenar siempre con
// Frenar(); la bicicleta playera es «código viejo» que solo sabe
// PedalearHaciaAtras(). PlayeraAdapter traduce una interfaz a la otra sin
// tocar el código legacy ni el cliente.

import "fmt"

// Bicicleta es la interfaz que espera el cliente: toda bicicleta sabe frenar.
type Bicicleta interface {
	Frenar()
}

// Playera es la bicicleta playera, el «código viejo»: solo sabe frenar
// pedaleando hacia atrás.
type Playera struct{}

// PedalearHaciaAtras es la única forma de frenar que conoce la playera.
func (p *Playera) PedalearHaciaAtras() {
	fmt.Println("Frenando la bicicleta playera...")
}

// BicicletaDeCarreras es la bicicleta moderna que frena con Frenar().
type BicicletaDeCarreras struct{}

// Frenar frena la bicicleta de carrera.
func (c *BicicletaDeCarreras) Frenar() {
	fmt.Println("Frenando la bicicleta de carrera...")
}

// PlayeraAdapter adapta la playera a la interfaz Bicicleta: expone Frenar()
// y traduce la llamada a PedalearHaciaAtras().
type PlayeraAdapter struct {
	Playera
}

// Frenar adapta el freno de la playera para que el cliente pueda usarla como
// cualquier otra bicicleta.
func (b *PlayeraAdapter) Frenar() {
	fmt.Println("Adaptando el freno de la bicicleta playera para la bicicleta de carrera...")
	fmt.Print(" -> ")
	b.PedalearHaciaAtras()
}
