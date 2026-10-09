package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// animarSeleccion muestra en w el estado del arreglo paso a paso mientras lo
// ordena con selección: una línea con el arreglo y una barra de # por
// elemento, esperando delay entre pasos. El mínimo recién colocado y la
// posición donde estaba quedan marcados con *. Trabaja sobre una copia, así
// que no modifica el arreglo original. No imprime por su cuenta: todo el
// dibujo sale por w.
func animarSeleccion(w io.Writer, arr []int, delay time.Duration) {
	copia := append([]int(nil), arr...)
	dibujarEstado(w, copia, "Estado inicial", 0, 0, delay)

	for i := 0; i < len(copia)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(copia); j++ {
			if copia[j] < copia[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			copia[i], copia[minIdx] = copia[minIdx], copia[i]
		}
		desde, hasta := i, i+1
		if minIdx < desde {
			desde = minIdx
		}
		if minIdx >= hasta {
			hasta = minIdx + 1
		}
		dibujarEstado(w, copia,
			fmt.Sprintf("Paso %d: coloca el mínimo %d en arr[%d]", i+1, copia[i], i),
			desde, hasta, delay)
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
