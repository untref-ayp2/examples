package main

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const anchoMaximoCamino = 60

// dibujarRecorrido escribe en w el camino de 0 a distancia con las estaciones
// en su posición kilométrica: S marca la partida, E una estación de servicio,
// * una estación donde se recarga y D el destino. Debajo se rotulan los
// kilómetros de cada estación. Espera delay al terminar. No imprime por su
// cuenta: todo el dibujo sale por w.
func dibujarRecorrido(w io.Writer, distancia int, estaciones []int, paradas []int, delay time.Duration) {
	// escala: un carácter por cada N kilómetros, para que el camino no se
	// estire de más en terminales angostas
	escala := 1
	if distancia > anchoMaximoCamino {
		escala = (distancia + anchoMaximoCamino - 1) / anchoMaximoCamino
	}

	ncols := distancia / escala
	if ncols < 1 {
		ncols = 1
	}

	esParada := make(map[int]bool, len(paradas))
	for _, p := range paradas {
		esParada[p] = true
	}

	camino := make([]byte, ncols+1)
	for i := range camino {
		camino[i] = '-'
	}
	camino[0] = 'S'
	for _, e := range estaciones {
		col := e / escala
		if col < 0 || col >= len(camino) {
			continue
		}
		if esParada[e] {
			camino[col] = '*'
		} else {
			camino[col] = 'E'
		}
	}
	camino[len(camino)-1] = 'D'

	fmt.Fprintf(w, "0 %s %d\n", string(camino), distancia)

	// etiquetas con los kilómetros de cada estación, debajo de su marcador
	etiquetas := make([]byte, len(camino)+2)
	for i := range etiquetas {
		etiquetas[i] = ' '
	}
	for _, e := range estaciones {
		texto := strconv.Itoa(e)
		col := 2 + e/escala // "0 " es el prefijo del camino
		ini := col - len(texto) + 1
		if ini < 0 {
			ini = 0
		}
		fin := ini + len(texto)
		if fin > len(etiquetas) {
			fin = len(etiquetas)
			texto = texto[len(texto)-(fin-ini):]
		}
		copy(etiquetas[ini:fin], texto)
	}
	fmt.Fprintln(w, string(etiquetas))

	fmt.Fprintln(w, "S: partida · E: estación · *: parada donde se recarga · D: destino")

	time.Sleep(delay)
}
