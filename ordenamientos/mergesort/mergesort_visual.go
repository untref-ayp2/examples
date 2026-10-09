package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// animarMergesort muestra en w el estado del arreglo paso a paso mientras lo
// ordena con mergesort: una línea con el arreglo y una barra de # por
// elemento, esperando delay entre pasos. El segmento que se acaba de
// fusionar queda marcado con *. Trabaja sobre una copia, así que no modifica
// el arreglo original. No imprime por su cuenta: todo el dibujo sale por w.
func animarMergesort(w io.Writer, array []int, delay time.Duration) {
	copia := append([]int(nil), array...)
	dibujarEstado(w, copia, "Estado inicial", 0, 0, delay)

	aux := make([]int, len(copia))
	mergesortVisual(w, copia, aux, 0, len(copia), delay)
}

// mergesortVisual ordena en su lugar el segmento copia[lo:hi] con mergesort
// y dibuja el arreglo completo después de cada fusión.
func mergesortVisual(w io.Writer, copia, aux []int, lo, hi int, delay time.Duration) {
	if hi-lo < 2 {
		return
	}
	medio := (lo + hi) / 2

	mergesortVisual(w, copia, aux, lo, medio, delay)
	mergesortVisual(w, copia, aux, medio, hi, delay)

	// intercala las dos mitades ordenadas en aux y las copia de vuelta
	i, j := lo, medio
	for k := lo; k < hi; k++ {
		if i < medio && (j >= hi || copia[i] <= copia[j]) {
			aux[k] = copia[i]
			i++
		} else {
			aux[k] = copia[j]
			j++
		}
	}
	copy(copia[lo:hi], aux[lo:hi])

	dibujarEstado(w, copia,
		fmt.Sprintf("Fusiona el segmento arr[%d:%d]", lo, hi),
		lo, hi, delay)
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

	ancho := 1
	for _, v := range arr {
		if a := len(strconv.Itoa(v)); a > ancho {
			ancho = a
		}
	}
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
