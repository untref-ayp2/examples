package main

import "strings"

// Este archivo contiene el backtracking del laberinto de la rata, sin
// animación: la búsqueda no notifica nada, solo devuelve si encontró camino.
//
// Es el mismo esquema que el de la versión animada (carpeta ratmaze/):
// esSolucion, poda (esFactible), registrar, extender y borrar. La única
// diferencia es que acá no existe el callback PasoLaberinto.

// RatMaze resuelve el laberinto de la rata: encontrar un camino desde la
// celda superior izquierda hasta la inferior derecha, moviéndose en las
// cuatro direcciones por las celdas libres (true).
type RatMaze struct {
	maze     [][]bool
	solucion [][]bool // celdas que forman el camino encontrado
	visitado [][]bool // celdas exploradas (incluye las abandonadas)
	destX    int
	destY    int
}

// NuevoRatMaze crea un resolutor para el laberinto dado.
//
// Precondición: maze no está vacío, todas sus filas tienen la misma
// cantidad de columnas y hay al menos una celda libre. El laberinto no se
// copia: se guarda la referencia.
func NuevoRatMaze(maze [][]bool) *RatMaze {
	return &RatMaze{maze: maze}
}

// Resolver busca un camino y deja marcadas en solucion las celdas que lo
// forman. Devuelve true si encontró alguno.
//
// Cada llamada reinicia el estado (solucion, visitado y destino), así que el
// mismo RatMaze se puede resolver varias veces. El camino encontrado depende
// del orden en que se prueban las direcciones (derecha, abajo, izquierda,
// arriba): no es necesariamente el más corto.
func (r *RatMaze) Resolver() bool {
	r.solucion = make([][]bool, len(r.maze))
	r.visitado = make([][]bool, len(r.maze))
	for i := range r.maze {
		r.solucion[i] = make([]bool, len(r.maze[i]))
		r.visitado[i] = make([]bool, len(r.maze[i]))
	}
	r.destY = len(r.maze) - 1
	r.destX = len(r.maze[r.destY]) - 1
	return r.backtracking(0, 0)
}

// esSolucion indica si (x, y) es el destino del laberinto (y es transitable).
func (r *RatMaze) esSolucion(x, y int) bool {
	return x == r.destX && y == r.destY && r.maze[y][x]
}

// esFactible indica si se puede avanzar a (x, y): está dentro del laberinto,
// es una celda libre y no forma parte del camino actual.
func (r *RatMaze) esFactible(x, y int) bool {
	return y >= 0 && y < len(r.maze) &&
		x >= 0 && x < len(r.maze[y]) &&
		r.maze[y][x] &&
		!r.solucion[y][x]
}

// registrar marca (x, y) como parte del camino.
func (r *RatMaze) registrar(x, y int) {
	r.solucion[y][x] = true
	r.visitado[y][x] = true
}

// borrar desmarca (x, y) al volver atrás.
func (r *RatMaze) borrar(x, y int) {
	r.solucion[y][x] = false
}

// backtracking es la recursión general del esquema: marca la celda actual y
// prueba las cuatro direcciones hasta llegar al destino.
//
// Termina porque cada llamada recursiva avanza a una celda que no estaba en
// el camino (esFactible exige !solucion[y][x]) y el laberinto es finito: no
// hay ciclos, así que ninguna rama reexplora la misma celda.
func (r *RatMaze) backtracking(x, y int) bool {
	// esSolucion: llegamos al destino.
	if r.esSolucion(x, y) {
		r.registrar(x, y)
		return true
	}

	// Poda: celda fuera del laberinto, con pared, o ya incluida en el camino.
	if !r.esFactible(x, y) {
		return false
	}

	// registrar: marco la celda como parte del camino.
	r.registrar(x, y)

	// extender: pruebo las cuatro direcciones. El orden define cuál de
	// todos los caminos posibles se devuelve.
	if r.backtracking(x+1, y) { // derecha
		return true
	}
	if r.backtracking(x, y+1) { // abajo
		return true
	}
	if r.backtracking(x-1, y) { // izquierda
		return true
	}
	if r.backtracking(x, y-1) { // arriba
		return true
	}

	// borrar (vuelta atrás): la celda no lleva a ninguna solución.
	r.borrar(x, y)
	return false
}

// Tablero devuelve una impresión estática del laberinto: las celdas del
// camino como 'o', las paredes como '#' y el resto como '.'. No hay cursor,
// colores ni animación: es la foto final de la búsqueda.
//
// Precondición: llamar después de Resolver (si no, solucion y visitado son
// nil y todas las celdas libres se ven como '.').
func (r *RatMaze) Tablero() string {
	var b strings.Builder
	for y := range r.maze {
		for x := range r.maze[y] {
			switch {
			case r.solucion != nil && r.solucion[y][x]:
				b.WriteString("o ")
			case r.maze[y][x]:
				b.WriteString(". ")
			default:
				b.WriteString("# ")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}
