package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// dibujarLineaTiempo escribe en w una línea de tiempo con las actividades
// seleccionadas: cada fila es una actividad y su barra va de Inicio a Fin.
// Espera delay al terminar. No imprime por su cuenta: todo el dibujo sale
// por w.
func dibujarLineaTiempo(w io.Writer, actividades []Actividad, delay time.Duration) {
	if len(actividades) == 0 {
		fmt.Fprintln(w, "No hay actividades seleccionadas.")
		time.Sleep(delay)
		return
	}

	// el eje de tiempos llega hasta la actividad que termina más tarde
	fin := 0
	anchoNombre := 0
	for _, a := range actividades {
		if a.Fin > fin {
			fin = a.Fin
		}
		if len(a.Nombre) > anchoNombre {
			anchoNombre = len(a.Nombre)
		}
	}

	// eje de tiempos (una celda de 2 caracteres por unidad de tiempo)
	fmt.Fprintf(w, "%-*s", anchoNombre+1, "")
	for t := 0; t < fin; t++ {
		fmt.Fprintf(w, "%2d", t)
	}
	fmt.Fprintln(w)

	// una fila por actividad, con su barra en las unidades que ocupa
	for _, a := range actividades {
		var barra strings.Builder
		for t := 0; t < fin; t++ {
			switch {
			case t < a.Inicio || t >= a.Fin:
				barra.WriteString("  ")
			case t == a.Inicio && t == a.Fin-1:
				barra.WriteString("[-]")
			case t == a.Inicio:
				barra.WriteString("[=")
			case t == a.Fin-1:
				barra.WriteString("=]")
			default:
				barra.WriteString("==")
			}
		}
		fmt.Fprintf(w, "%-*s %s\n", anchoNombre, a.Nombre, barra.String())
	}

	time.Sleep(delay)
}
