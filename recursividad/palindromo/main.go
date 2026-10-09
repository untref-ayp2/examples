package main

import "fmt"

func main() {
	palabras := []string{
		"reconocer",
		"NEUQUEN",
		"DABALEARROZALAZORRAELABAD",
		"Holanda",
		"ANITA LAVA LA TINA",
	}

	for _, palabra := range palabras {
		fmt.Printf("EsPalindromo(%q) = %v\n", palabra, EsPalindromo(palabra))
	}

	fmt.Println()
	fmt.Println("Ojo: compara los caracteres exactos, así que distingue mayúsculas")
	fmt.Println("y no ignora espacios ni acentos.")
}
