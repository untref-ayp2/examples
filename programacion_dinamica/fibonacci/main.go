package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	const delay = 200 * time.Millisecond

	dibujarArbolLlamadas(os.Stdout, 5, delay)
	fmt.Println()
	dibujarTablaFibonacci(os.Stdout, 8, delay)
}
