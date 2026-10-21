package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// animarInsercion muestra en w el estado del arreglo paso a paso mientras lo
// ordena con inserción: una línea con el arreglo y una barra de # por
// elemento, esperando delay entre pasos. Los elementos que se acaban de
// desplazar e insertar quedan marcados con *. Trabaja sobre una copia, así
// que no modifica el arreglo original. No imprime por su cuenta: todo el
// dibujo sale por w.
func animarInsercion(w io.Writer, arr []int, delay time.Duration) {
	copia := append([]int(nil), arr...)
	dibujarEstado(w, copia, "Estado inicial", 0, 0, delay)

	for i := 1; i < len(copia); i++ {
		clave := copia[i]
		j := i - 1
		for j >= 0 && copia[j] > clave {
			copia[j+1] = copia[j]
			j--
		}
		copia[j+1] = clave
		pos := j + 1
		dibujarEstado(w, copia,
			fmt.Sprintf("Inserta %d en la posición %d", clave, pos),
			pos, i+1, delay)
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
