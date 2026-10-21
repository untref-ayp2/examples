package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// animarQuicksort muestra en w el estado del arreglo paso a paso mientras lo
// ordena con quicksort: una línea con el arreglo y una barra de # por
// elemento, esperando delay entre pasos. La posición donde queda el pivote
// se marca con *. Trabaja sobre una copia, así que no modifica el arreglo
// original. No imprime por su cuenta: todo el dibujo sale por w.
func animarQuicksort(w io.Writer, array []int, delay time.Duration) {
	copia := append([]int(nil), array...)
	dibujarEstado(w, copia, "Estado inicial", 0, 0, delay)

	quicksortVisual(w, copia, 0, len(copia), delay)
}

// quicksortVisual particiona en su lugar el segmento copia[lo:hi] con
// quicksort y dibuja el arreglo completo después de cada partición, cuando
// el pivote queda en su posición final.
func quicksortVisual(w io.Writer, copia []int, lo, hi int, delay time.Duration) {
	if hi-lo < 2 {
		return
	}
	pivote := copia[lo]
	i := lo + 1 // borde: todo lo que está antes de i es menor que el pivote
	for j := lo + 1; j < hi; j++ {
		if copia[j] < pivote {
			copia[i], copia[j] = copia[j], copia[i]
			i++
		}
	}
	// el pivote va al borde de los menores y queda en su posición final
	piv := i - 1
	copia[lo], copia[piv] = copia[piv], copia[lo]

	dibujarEstado(w, copia,
		fmt.Sprintf("Particiona arr[%d:%d] con pivote %d (queda en arr[%d])", lo, hi, pivote, piv),
		piv, piv+1, delay)

	quicksortVisual(w, copia, lo, piv, delay)
	quicksortVisual(w, copia, piv+1, hi, delay)
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
