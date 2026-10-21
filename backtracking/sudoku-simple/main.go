package main

import "fmt"

func main() {
	// Tablero de ejemplo con solución única. Los 0 son celdas vacías. Para
	// experimentar, sacá alguna pista y mirá cuántas soluciones aparecen.
	tablero := [9][9]int{
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

	sudoku := NuevoSudoku(tablero)

	if !sudoku.Resolver() {
		fmt.Println("El tablero no tiene solución.")
		return
	}

	fmt.Printf("Tablero resuelto (%d solución):\n", len(sudoku.soluciones))
	fmt.Print(sudoku.Tablero())
}
