package main

// Este archivo contiene el backtracking del problema de las N reinas, sin
// animación.
//
// La solución parcial es un []int: la posición i corresponde a la fila i y su
// valor es la columna donde va la reina de esa fila. El esquema es el mismo
// que el de la versión animada (carpeta reinas/): esSolucion, esFactible,
// registrar, borrar y extender.

// NReinas resuelve el problema de las N reinas usando backtracking.
// Devuelve un slice con la columna de la reina de cada fila (el índice es la
// fila) o nil si no hay ninguna solución.
//
// Precondición: n >= 1.
func NReinas(n int) []int {
	var soluciones [][]int
	backtracking(n, 0, nil, &soluciones, true)
	if len(soluciones) == 0 {
		return nil
	}
	return soluciones[0]
}

// NReinasTodas devuelve todas las soluciones del problema de las N reinas.
// Cada solución es un slice de n columnas (una por fila). Con n = 2 y n = 3
// devuelve un slice vacío, porque no hay forma de ubicarlas.
//
// Precondición: n >= 1.
func NReinasTodas(n int) [][]int {
	var soluciones [][]int
	backtracking(n, 0, nil, &soluciones, false)
	return soluciones
}

// esSolucion verifica si la solución parcial ya tiene una reina en cada fila.
func esSolucion(n int, solucionParcial []int) bool {
	return len(solucionParcial) == n
}

// esFactible verifica si se puede colocar una reina en (fila, columna) dada
// una configuración parcial de reinas ya colocadas.
//
// Una reina en (fila, columna) ataca a la de (i, parcial[i]) si comparten
// columna (parcial[i] == columna) o alguna diagonal. En las diagonales, la
// suma fila+columna es constante en una dirección y la resta fila-columna en
// la otra: por eso alcanza con comparar esas dos cuentas.
func esFactible(fila int, columna int, solucionParcial []int) bool {
	for i := range solucionParcial {
		if solucionParcial[i] == columna ||
			fila+columna == i+solucionParcial[i] ||
			fila-columna == i-solucionParcial[i] {
			return false
		}
	}
	return true
}

// registrar agrega la columna elegida a la solución parcial.
func registrar(solucionParcial []int, columna int) []int {
	return append(solucionParcial, columna)
}

// borrar quita la última reina registrada (vuelta atrás).
func borrar(solucionParcial []int) []int {
	return solucionParcial[:len(solucionParcial)-1]
}

// backtracking es la recursión general del esquema: coloca una reina por fila
// y prueba todas las columnas de la fila actual.
//
// Si primera es true, corta al encontrar la primera solución; si es false,
// sigue hasta agotar todas y las acumula en soluciones.
//
// Termina porque cada llamada recursiva baja a la fila siguiente
// (fila+1) y el tablero tiene n filas: a lo sumo hay n niveles de recursión,
// y el bucle de columnas es finito.
func backtracking(n int, fila int, solucionParcial []int, soluciones *[][]int, primera bool) {
	// Caso base (esSolucion): ya hay una reina colocada en cada fila.
	if esSolucion(n, solucionParcial) {
		// Copiar es obligatorio: solucionParcial se sigue modificando con
		// registrar/borrar y, sin copia, todas las soluciones guardadas
		// terminarían apuntando al mismo arreglo.
		copia := make([]int, len(solucionParcial))
		copy(copia, solucionParcial)
		*soluciones = append(*soluciones, copia)
		return
	}

	// Extender: probar cada columna de la fila actual.
	for columna := 0; columna < n; columna++ {
		// esFactible: la reina no entra en conflicto con las ya colocadas.
		if !esFactible(fila, columna, solucionParcial) {
			continue
		}

		// Registrar: colocar la reina en (fila, columna).
		solucionParcial = registrar(solucionParcial, columna)

		// Llamada recursiva: pasar a la siguiente fila.
		backtracking(n, fila+1, solucionParcial, soluciones, primera)

		// Borrar (vuelta atrás): quitar la reina que se puso en esta fila.
		solucionParcial = borrar(solucionParcial)

		// Con primera=true ya hay una solución guardada: no hace falta
		// seguir probando columnas.
		if primera && len(*soluciones) > 0 {
			return
		}
	}

	// Ninguna columna funcionó en esta fila: vuelta atrás.
}
