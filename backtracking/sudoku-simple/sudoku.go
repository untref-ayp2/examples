package main

import "strings"

// Este archivo contiene el backtracking del sudoku, sin animación.
//
// Es el mismo esquema que el de la versión animada (carpeta sudoku/):
// esSolucion, proximaVacia, esFactible, registrar, borrar, guardarSolucion y
// extender. La diferencia es que acá el backtracking no recibe ningún
// callback.

// Sudoku resuelve un tablero 9×9 con backtracking. Las celdas vacías se
// representan con 0. Una vez resuelto guarda cada tablero completo en
// soluciones.
type Sudoku struct {
	tablero    [9][9]int
	soluciones [][9][9]int
}

// NuevoSudoku crea un resolutor para el tablero dado. El tablero se copia:
// el original no se modifica.
//
// Precondición: cada celda es 0 (vacía) o un valor entre 1 y 9; el tablero
// puede ser contradictorio, en ese caso Resolver devuelve false.
func NuevoSudoku(tablero [9][9]int) *Sudoku {
	return &Sudoku{tablero: tablero}
}

// Resolver busca todas las soluciones del tablero y las deja en soluciones.
// Devuelve true si existe al menos una. El tablero que se pasó a NuevoSudoku
// no se modifica, y al terminar el tablero interno queda con la primera
// solución, para que Tablero la pueda mostrar.
func (s *Sudoku) Resolver() bool {
	s.soluciones = nil
	// backtracking devuelve true solo cuando guarda una solución, pero
	// sigue explorando el resto: para saber si había alguna, alcanza con
	// mirar si se guardó al menos una.
	s.backtracking()
	// Cuando la búsqueda termina, el tablero volvió a su estado inicial:
	// todas las asignaciones se deshicieron al retroceder. Si hubo solución,
	// se copia la primera de vuelta para que Tablero la muestre.
	if len(s.soluciones) > 0 {
		s.tablero = s.soluciones[0]
	}
	return len(s.soluciones) > 0
}

// esSolucion verifica si no queda ninguna celda vacía en el tablero.
func (s *Sudoku) esSolucion() bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if s.tablero[y][x] == 0 {
				return false
			}
		}
	}
	return true
}

// proximaVacia devuelve la primera celda vacía, es decir, la próxima a
// extender. Devuelve (-1, -1) si no hay ninguna.
func (s *Sudoku) proximaVacia() (x, y int) {
	for y = 0; y < 9; y++ {
		for x = 0; x < 9; x++ {
			if s.tablero[y][x] == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

// esFactible verifica si el valor n puede ir en (x, y) sin repetirse en su
// fila, su columna ni su cuadrante de 3×3.
func (s *Sudoku) esFactible(x, y, n int) bool {
	for i := 0; i < 9; i++ {
		if s.tablero[y][i] == n || s.tablero[i][x] == n {
			return false
		}
	}
	// El cuadrante que contiene a (x, y) arranca en la celda (x0, y0).
	y0 := (y / 3) * 3
	x0 := (x / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s.tablero[y0+i][x0+j] == n {
				return false
			}
		}
	}
	return true
}

// registrar coloca el valor n en la celda (x, y).
func (s *Sudoku) registrar(x, y, n int) {
	s.tablero[y][x] = n
}

// borrar quita el valor de la celda (x, y) (vuelta atrás).
func (s *Sudoku) borrar(x, y int) {
	s.tablero[y][x] = 0
}

// guardarSolucion guarda una copia del tablero completo. La copia es
// necesaria: el tablero se sigue modificando con registrar y borrar.
func (s *Sudoku) guardarSolucion() {
	var copia [9][9]int
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			copia[y][x] = s.tablero[y][x]
		}
	}
	s.soluciones = append(s.soluciones, copia)
}

// backtracking es la recursión general del esquema: busca la próxima celda
// vacía y prueba los valores 1 a 9.
//
// Termina porque cada llamada recursiva llena una celda más: hay a lo sumo 81
// niveles de recursión y el bucle de valores es finito. Cuando una rama no
// tiene salida, sus casillas se deshacen y la recursión retrocede.
func (s *Sudoku) backtracking() bool {
	// Caso base (esSolucion): no quedan celdas vacías.
	if s.esSolucion() {
		s.guardarSolucion()
		return true
	}

	// Extender: probar los valores posibles en la próxima celda vacía.
	x, y := s.proximaVacia()
	for n := 1; n <= 9; n++ {
		if !s.esFactible(x, y, n) {
			continue
		}

		// Registrar: probar el valor n en la celda vacía.
		s.registrar(x, y, n)

		// Llamada recursiva: resolver el resto del tablero.
		s.backtracking()

		// Borrar (vuelta atrás): deshacer la asignación.
		s.borrar(x, y)
	}

	// Ningún valor sirve: esta rama no tiene solución.
	return false
}

// Tablero devuelve una representación del tablero con bordes: los dígitos en
// su lugar y las celdas vacías como 0.
//
// Después de Resolver muestra la primera solución encontrada; antes, el
// tablero inicial.
func (s *Sudoku) Tablero() string {
	var b strings.Builder
	b.WriteString("┏━━━━━━━┳━━━━━━━┳━━━━━━━┓\n")
	for y := 0; y < 9; y++ {
		b.WriteString("┃ ")
		for x := 0; x < 9; x++ {
			if x == 3 || x == 6 {
				b.WriteString("┃ ")
			}
			b.WriteString(string(rune('0' + s.tablero[y][x])))
			b.WriteString(" ")
			if x == 8 {
				b.WriteString("┃")
			}
		}
		if y == 2 || y == 5 {
			b.WriteString("\n┣━━━━━━━╋━━━━━━━╋━━━━━━━┫\n")
		} else {
			b.WriteString("\n")
		}
	}
	b.WriteString("┗━━━━━━━┻━━━━━━━┻━━━━━━━┛\n")
	return b.String()
}
