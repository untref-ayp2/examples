package main

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// dibujarTablaMochila escribe en w la tabla DP de la mochila 0/1: una fila por
// cada cantidad de ítems considerados (i) y una columna por capacidad (c),
// con el máximo valor alcanzable en cada celda. Espera delay entre fila y
// fila. No imprime por su cuenta: todo el dibujo sale por w.
func dibujarTablaMochila(w io.Writer, items []Item, capacidad int, delay time.Duration) {
	n := len(items)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacidad+1)
	}
	for i := 1; i <= n; i++ {
		for c := 1; c <= capacidad; c++ {
			if items[i-1].Peso > c {
				dp[i][c] = dp[i-1][c]
			} else if con := items[i-1].Valor + dp[i-1][c-items[i-1].Peso]; con > dp[i-1][c] {
				dp[i][c] = con
			} else {
				dp[i][c] = dp[i-1][c]
			}
		}
	}

	ancho := 1
	for i := 0; i <= n; i++ {
		for c := 0; c <= capacidad; c++ {
			if a := len(strconv.Itoa(dp[i][c])); a > ancho {
				ancho = a
			}
		}
	}
	celda := fmt.Sprintf(" %%%dd", ancho)

	fmt.Fprintf(w, "Mochila 0/1: %d ítems, capacidad %d\n\n", n, capacidad)
	fmt.Fprintf(w, "%8s", "cap.")
	for c := 0; c <= capacidad; c++ {
		fmt.Fprintf(w, celda, c)
	}
	fmt.Fprintln(w)
	for i := 0; i <= n; i++ {
		fmt.Fprintf(w, "%8s", fmt.Sprintf("i=%d", i))
		for c := 0; c <= capacidad; c++ {
			fmt.Fprintf(w, celda, dp[i][c])
		}
		fmt.Fprintln(w)
		time.Sleep(delay)
	}
	fmt.Fprintf(w, "\nValor máximo: %d\n", dp[n][capacidad])
}
