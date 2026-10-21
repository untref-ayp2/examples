package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// animarBurbujeo muestra en w el estado del arreglo paso a paso mientras lo
// ordena con burbujeo: una línea con el arreglo y una barra de # por
// elemento, esperando delay entre pasos. Los dos elementos que se acaban de
// intercambiar quedan marcados con *. Trabaja sobre una copia, así que no
// modifica el arreglo original. No imprime por su cuenta: todo el dibujo
// sale por w.
func animarBurbujeo(w io.Writer, arr []int, delay time.Duration) {
	copia := append([]int(nil), arr...)
	dibujarEstado(w, copia, "Estado inicial", 0, 0, delay)

	for i := 0; i < len(copia)-1; i++ {
		intercambiado := false
		for j := 0; j < len(copia)-i-1; j++ {
			if copia[j] > copia[j+1] {
				a, b := copia[j], copia[j+1]
				copia[j], copia[j+1] = copia[j+1], copia[j]
				intercambiado = true
				dibujarEstado(w, copia,
					fmt.Sprintf("Pasada %d: intercambia arr[%d]=%d con arr[%d]=%d", i+1, j, a, j+1, b),
					j, j+2, delay)
			}
		}
		if !intercambiado {
			break
		}
	}
}

// dibujarEstado escribe en w el estado actual de arr: una línea con el
// arreglo y una barra de # por elemento, proporcional a su valor. Las
// posiciones en [desde, hasta) se marcan con * a la derecha. Espera delay
// al terminar. No imprime por su cuenta.
func dibujarEstado(w io.Writer, arr []int, titulo string, desde, hasta int, delay time.Duration) {
	fmt.Fprintln(w, titulo)

	fmt.Fprint(w, "[")
	for i, v := range arr {
		if i > 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, v)
	}
	fmt.Fprintln(w, "]")

	// ancho es la cantidad de dígitos del número más largo, para que la
	// columna de valores quede alineada con %*d (ancho variable).
	ancho := 1
	for _, v := range arr {
		if a := len(strconv.Itoa(v)); a > ancho {
			ancho = a
		}
	}
	// Cada elemento se dibuja como una barra de # proporcional a su valor.
	// Precondición: los elementos son >= 0 (strings.Repeat con una cantidad
	// negativa paniquea).
	for i, v := range arr {
		marca := " "
		if i >= desde && i < hasta {
			marca = "*"
		}
		fmt.Fprintf(w, "%*d |%s %s\n", ancho, v, strings.Repeat("#", v), marca)
	}
	fmt.Fprintln(w)

	time.Sleep(delay)
}
