package main

import (
	"fmt"
	"time"
)

func main() {
	// para experimentar: subí n de a poco y mirá cómo explota el tiempo
	// (la variante recursiva es exponencial)
	for _, n := range []int{10, 20, 30, 35} {
		start := time.Now()
		resultado := Fibonacci(n)
		fmt.Printf("Fibonacci(%d) = %d (%s)\n", n, resultado, time.Since(start))
	}
}
