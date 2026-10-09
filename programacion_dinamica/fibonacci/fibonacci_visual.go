package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// dibujarArbolLlamadas escribe en w el árbol de llamadas de Fibonacci(n),
// rotulando cada nodo con su resultado. Espera delay entre nodo y nodo.
// No imprime por su cuenta: todo el dibujo sale por w.
func dibujarArbolLlamadas(w io.Writer, n int, delay time.Duration) {
	valores := make([]int, n+1)
	valores[0] = 0
	if n >= 1 {
		valores[1] = 1
	}
	for i := 2; i <= n; i++ {
		valores[i] = valores[i-1] + valores[i-2]
	}

	fmt.Fprintf(w, "Árbol de llamadas de Fibonacci(%d):\n", n)
	dibujarNodo(w, n, 0, valores, delay)
}

func dibujarNodo(w io.Writer, n, nivel int, valores []int, delay time.Duration) {
	indent := strings.Repeat("  ", nivel)
	fmt.Fprintf(w, "%sfib(%d) = %d\n", indent, n, valores[n])
	time.Sleep(delay)
	if n >= 2 {
		dibujarNodo(w, n-1, nivel+1, valores, delay)
		dibujarNodo(w, n-2, nivel+1, valores, delay)
	}
}

// dibujarTablaFibonacci escribe en w la tabla bottom-up dp[i] = Fibonacci(i)
// fila por fila, con un delay entre cada fila. No imprime por su cuenta:
// todo el dibujo sale por w.
func dibujarTablaFibonacci(w io.Writer, n int, delay time.Duration) {
	dp := make([]int, n+1)
	fmt.Fprintln(w, "Tabla bottom-up: dp[i] = Fibonacci(i)")
	fmt.Fprintf(w, "dp[0] = 0\n")
	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[1] = 1
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
		fmt.Fprintf(w, "dp[%d] = %d\n", i, dp[i])
		time.Sleep(delay)
	}
}
