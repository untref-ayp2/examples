package main

import "fmt"

// main muestra el valor de n! para n de 0 a 20 usando Factorial.
func main() {
	for n := 0; n <= 20; n++ {
		fmt.Printf("%d! = %d\n", n, Factorial(n))
	}
}
