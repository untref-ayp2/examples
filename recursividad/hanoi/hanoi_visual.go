package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// estadoHanoi representa el estado de las tres varas. Cada vara guarda sus
// discos del más chico (arriba) al más grande (abajo).
type estadoHanoi struct {
	discos int
	varas  [3][]int
}

// nuevoEstadoHanoi crea un estado inicial con los n discos apilados en la
// primera vara.
func nuevoEstadoHanoi(n int) *estadoHanoi {
	estado := &estadoHanoi{discos: n}
	for disco := n; disco >= 1; disco-- {
		estado.varas[0] = append(estado.varas[0], disco)
	}
	return estado
}

// mover pasa el disco superior de la vara origen a la vara destino.
func (e *estadoHanoi) mover(disco, origen, destino int) {
	tope := e.varas[origen][len(e.varas[origen])-1]
	if tope != disco {
		panic("se intentó mover un disco que no está en el tope")
	}
	e.varas[origen] = e.varas[origen][:len(e.varas[origen])-1]
	e.varas[destino] = append(e.varas[destino], disco)
}

func (e *estadoHanoi) alturaMaxima() int {
	return e.discos
}

// dibujar escribe el estado actual de las tres varas en w y espera delay
// antes de devolver. No imprime por su cuenta: todo el dibujo sale por w.
func (e *estadoHanoi) dibujar(w io.Writer, delay time.Duration) {
	const vara = "|"
	const disco = "#"

	alto := e.alturaMaxima()
	anchoTotal := 2*alto + 1

	centrar := func(contenido string) string {
		espacios := (anchoTotal - len(contenido)) / 2
		return strings.Repeat(" ", espacios) + contenido + strings.Repeat(" ", anchoTotal-len(contenido)-espacios)
	}

	for nivel := alto; nivel >= 1; nivel-- {
		for _, varaDiscos := range e.varas {
			if len(varaDiscos) < nivel {
				fmt.Fprint(w, centrar(vara))
				continue
			}
			discoTope := varaDiscos[len(varaDiscos)-nivel]
			fmt.Fprint(w, centrar(strings.Repeat(disco, 2*discoTope+1)))
		}
		fmt.Fprintln(w)
	}

	for _, nombre := range []string{"A", "B", "C"} {
		fmt.Fprint(w, centrar(nombre))
	}
	fmt.Fprintln(w)

	time.Sleep(delay)
}
