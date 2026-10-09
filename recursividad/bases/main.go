package main

import "fmt"

func main() {
	fmt.Println("ImprimirEnDecimal(127):", ImprimirEnDecimal(127))
	fmt.Println("ImprimirEnDecimal(2026):", ImprimirEnDecimal(2026))

	fmt.Println("ImprimirEnBase(42, 2):", ImprimirEnBase(42, 2))
	fmt.Println("ImprimirEnBase(255, 8):", ImprimirEnBase(255, 8))
	fmt.Println("ImprimirEnBase(255, 16):", ImprimirEnBase(255, 16))
	fmt.Println("ImprimirEnBase(2026, 16):", ImprimirEnBase(2026, 16))
}
