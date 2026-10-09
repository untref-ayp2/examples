package main

import "testing"

func TestPalindromos(t *testing.T) {
	palabras := []string{"NEUQUEN", "DABALEARROZALAZORRAELABAD", "reconocer"}
	for _, palabra := range palabras {
		if !EsPalindromo(palabra) {
			t.Errorf("EsPalindromo(%q) = false; se esperaba true", palabra)
		}
	}
}

func TestNoPalindromos(t *testing.T) {
	palabras := []string{"Holanda", "estudiar", "ANITA LAVA LA TINA"}
	for _, palabra := range palabras {
		if EsPalindromo(palabra) {
			t.Errorf("EsPalindromo(%q) = true; se esperaba false", palabra)
		}
	}
	// ¿cómo podemos modificar el código para que reconozca "ANITA LAVA LA TINA"?
}
