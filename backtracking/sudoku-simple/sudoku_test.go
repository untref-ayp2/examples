package main

import (
	"strings"
	"testing"
)

// tableroEjemplo es el mismo tablero del main de la versión animada: tiene
// solución única.
func tableroEjemplo() [9][9]int {
	return [9][9]int{
		{4, 0, 0, 1, 9, 5, 0, 6, 8},
		{8, 0, 0, 0, 0, 0, 7, 0, 1},
		{9, 0, 1, 6, 0, 0, 0, 3, 0},
		{0, 0, 7, 0, 2, 6, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0, 0, 3},
		{0, 1, 0, 8, 7, 0, 4, 0, 0},
		{0, 3, 0, 0, 0, 0, 8, 0, 5},
		{1, 0, 5, 0, 0, 0, 0, 0, 0},
		{7, 9, 0, 4, 0, 1, 0, 0, 0},
	}
}

// tableroResuelto es un sudoku completo y válido, con una sola celda vacía
// que se llena sola.
func tableroResuelto() [9][9]int {
	return [9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 0, 7, 9},
	}
}

// TestSudokuResuelve comprueba que Resolver devuelva true, que guarde la
// cantidad exacta de soluciones del tablero y que la primera sea una solución
// válida del original.
//
// El tablero del ejemplo tiene solución única: si la búsqueda cortara en la
// primera rama que sirve, este test no lo notaría (seguiría dando 1), pero sí
// lo notarían TestSudokuVariasSoluciones y TestSudokuTodasLasSoluciones.
func TestSudokuResuelve(t *testing.T) {
	original := tableroEjemplo()
	sudoku := NuevoSudoku(original)

	if !sudoku.Resolver() {
		t.Fatal("Resolver() = false, se esperaba encontrar la solución")
	}
	if len(sudoku.soluciones) != 1 {
		t.Fatalf("Resolver() guardó %d soluciones, se esperaba exactamente 1",
			len(sudoku.soluciones))
	}
	solucion := sudoku.soluciones[0]
	if !esSolucionValida(original, solucion) {
		t.Errorf("la solución no es válida:\n%s", sudoku.Tablero())
	}
}

// TestSudokuVariasSoluciones comprueba que, cuando el tablero admite más de
// una solución, Resolver las encuentre todas (y no sólo la primera).
//
// El tablero es el de la versión animada (carpeta sudoku/), que verifica la
// misma cantidad: así las dos versiones del ejemplo quedan atadas al mismo
// resultado conocido.
func TestSudokuVariasSoluciones(t *testing.T) {
	tablero := [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 0, 0},
	}

	sudoku := NuevoSudoku(tablero)
	if !sudoku.Resolver() {
		t.Fatal("Resolver() = false, el tablero tiene soluciones")
	}
	if len(sudoku.soluciones) != 2 {
		t.Fatalf("Resolver() guardó %d soluciones, se esperaban exactamente 2",
			len(sudoku.soluciones))
	}
	for i := range sudoku.soluciones {
		if !esSolucionValida(tablero, sudoku.soluciones[i]) {
			t.Errorf("la solución %d no es válida", i+1)
		}
	}
}

// TestSudokuTodasLasSoluciones comprueba la cantidad exacta de soluciones de
// un tablero con muchas: el corazón del ejemplo es que Resolver busca TODAS,
// no la primera. Un mutante que cortara al encontrar una solución daría 1, y
// uno que guardara cada solución dos veces daría 65208; con la cantidad
// exacta los dos fallan.
//
// El 32604 lo verifica también la versión animada (carpeta sudoku/), así que
// es la cifra de referencia de las dos.
func TestSudokuTodasLasSoluciones(t *testing.T) {
	tablero := [9][9]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 0, 0, 0, 4, 5, 6},
		{2, 3, 1, 0, 0, 0, 0, 0, 0},
		{5, 6, 4, 0, 0, 0, 0, 0, 0},
		{8, 9, 7, 0, 0, 0, 0, 0, 0},
		{3, 1, 2, 0, 0, 0, 0, 0, 0},
		{6, 4, 5, 0, 0, 0, 0, 0, 0},
		{9, 7, 8, 0, 0, 0, 0, 0, 0},
	}

	sudoku := NuevoSudoku(tablero)
	if !sudoku.Resolver() {
		t.Fatal("Resolver() = false, el tablero tiene soluciones")
	}
	if len(sudoku.soluciones) != 32604 {
		t.Fatalf("Resolver() guardó %d soluciones, se esperaban exactamente 32604",
			len(sudoku.soluciones))
	}
}

// TestSudokuSinSolucion comprueba que Resolver devuelva false cuando el
// tablero no tiene ninguna solución.
func TestSudokuSinSolucion(t *testing.T) {
	// Tablero tomado de la versión animada: en la esquina superior derecha
	// no puede ir ningún dígito, así que hay que agotar la búsqueda para
	// descubrirlo. Es IMPORTANTE que el caso sea de falla rápida: un
	// tablero contradictorio pero casi vacío tarda muchísimo, porque recién
	// se descubre al agotar el árbol entero.
	sinSolucion := [9][9]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
	}

	sudoku := NuevoSudoku(sinSolucion)

	if sudoku.Resolver() {
		t.Error("Resolver() = true, un tablero con una celda sin valores posibles no tiene solución")
	}
	if len(sudoku.soluciones) != 0 {
		t.Errorf("no debería haber soluciones guardadas, hay %d", len(sudoku.soluciones))
	}
}

// TestSudokuNoModificaElOriginal cubre dos cosas distintas.
//
// La primera (que el tablero del llamador no cambie) es una garantía del tipo
// de valor, no de la implementación: NuevoSudoku recibe [9][9]int por valor,
// así que ningún resolutor puede tocar la variable del llamador. La
// comparación se hace contra un literal con las pistas originales (no contra
// una copia hecha con :=, que sería siempre igual) para dejar el chequeo
// explícito, pero conviene saber que no puede fallar por un bug de
// implementación.
//
// La segunda sí tiene poder de detección: el tablero INTERNO del resolutor
// tiene que haber cambiado y coincidir con la primera solución guardada.
// Resolver asigna s.tablero = s.soluciones[0] al terminar justamente para que
// Tablero() muestre la solución; si esa asignación faltara, esta aserción
// falla. También se comprueba que dos instancias distintas con el mismo input
// no se pisen entre sí: si hubiera estado compartido (paquete o globales), la
// segunda cambiaría el resultado.
func TestSudokuNoModificaElOriginal(t *testing.T) {
	original := tableroEjemplo()
	esperado := [9][9]int{
		{4, 0, 0, 1, 9, 5, 0, 6, 8},
		{8, 0, 0, 0, 0, 0, 7, 0, 1},
		{9, 0, 1, 6, 0, 0, 0, 3, 0},
		{0, 0, 7, 0, 2, 6, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0, 0, 3},
		{0, 1, 0, 8, 7, 0, 4, 0, 0},
		{0, 3, 0, 0, 0, 0, 8, 0, 5},
		{1, 0, 5, 0, 0, 0, 0, 0, 0},
		{7, 9, 0, 4, 0, 1, 0, 0, 0},
	}

	sudoku := NuevoSudoku(original)
	if !sudoku.Resolver() {
		t.Fatal("Resolver() = false, el tablero tiene solución")
	}

	// Garantía por tipo de valor: la variable del llamador queda intacta.
	if original != esperado {
		t.Error("el tablero original no debería modificarse")
	}

	// Aserción con poder real: el tablero interno sí cambió y quedó igual a
	// la primera solución.
	if sudoku.tablero != sudoku.soluciones[0] {
		t.Error("el tablero interno debería quedar con la primera solución " +
			"después de Resolver")
	}
	if sudoku.tablero == esperado {
		t.Error("el tablero interno debería estar resuelto, no con las pistas originales")
	}

	// Dos instancias independientes con el mismo input no se pisan: si
	// compartieran estado, la segunda no daría lo mismo que la primera.
	otro := NuevoSudoku(original)
	if !otro.Resolver() {
		t.Fatal("Resolver() = false en la segunda instancia")
	}
	if len(otro.soluciones) != len(sudoku.soluciones) {
		t.Errorf("dos instancias con el mismo input dieron %d y %d soluciones",
			len(sudoku.soluciones), len(otro.soluciones))
	}
	if otro.soluciones[0] != sudoku.soluciones[0] {
		t.Error("dos instancias con el mismo input deberían encontrar la misma primera solución")
	}
}

// TestSudokuTablero comprueba que, después de resolver, la impresión muestre
// el tablero completo con bordes.
func TestSudokuTablero(t *testing.T) {
	sudoku := NuevoSudoku(tableroResuelto())
	sudoku.Resolver()

	tablero := sudoku.Tablero()
	if !strings.Contains(tablero, "┏") || !strings.Contains(tablero, "┛") {
		t.Errorf("el tablero debería dibujar bordes:\n%s", tablero)
	}
	if strings.Contains(tablero, "0") {
		t.Errorf("el tablero resuelto no debería tener celdas vacías:\n%s", tablero)
	}
}

// esSolucionValida comprueba que solucion sea un sudoku válido (cada fila,
// columna y cuadrante tiene los dígitos 1..9 sin repetir) y que respete las
// pistas del tablero original.
func esSolucionValida(original, solucion [9][9]int) bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if solucion[y][x] < 1 || solucion[y][x] > 9 {
				return false
			}
			if original[y][x] != 0 && original[y][x] != solucion[y][x] {
				return false
			}
		}
	}

	for i := 0; i < 9; i++ {
		var vistosFila, vistosColumna [10]bool
		for j := 0; j < 9; j++ {
			if vistosFila[solucion[i][j]] || vistosColumna[solucion[j][i]] {
				return false
			}
			vistosFila[solucion[i][j]] = true
			vistosColumna[solucion[j][i]] = true
		}
	}

	for caja := 0; caja < 9; caja++ {
		var vistos [10]bool
		y0 := (caja / 3) * 3
		x0 := (caja % 3) * 3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if vistos[solucion[y0+i][x0+j]] {
					return false
				}
				vistos[solucion[y0+i][x0+j]] = true
			}
		}
	}
	return true
}
