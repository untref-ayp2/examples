package main

import (
	"os"
	"time"
)

func main() {
	// para experimentar: cambiá la cantidad y las denominaciones
	denominaciones := []int{1, 2, 9, 10, 20, 50}
	AnimarCambio(27, denominaciones, os.Stdout, 30*time.Millisecond)
}
