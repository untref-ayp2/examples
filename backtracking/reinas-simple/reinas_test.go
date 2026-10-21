package main

import "testing"

// TestNReinasGeneraSolucion comprueba que NReinas devuelva una solución
// válida de n reinas para los n que tienen solución.
func TestNReinasGeneraSolucion(t *testing.T) {
	for _, n := range []int{1, 4, 8} {
		sol := NReinas(n)
		if len(sol) != n {
			t.Fatalf("NReinas(%d) = %v, se esperaban %d filas", n, sol, n)
		}
		if !esSolucionValida(sol) {
			t.Errorf("NReinas(%d) = %v no es una solución válida", n, sol)
		}
	}
}

// TestNReinasSinSolucion comprueba que devuelva nil cuando no hay solución
// (con 2 y 3 reinas no se puede ubicar ninguna sin que se ataquen).
func TestNReinasSinSolucion(t *testing.T) {
	for _, n := range []int{2, 3} {
		if sol := NReinas(n); sol != nil {
			t.Errorf("NReinas(%d) = %v, se esperaba nil", n, sol)
		}
	}
}

// TestNReinasTodas comprueba la cantidad de soluciones distintas de cada n.
func TestNReinasTodas(t *testing.T) {
	casos := []struct {
		n      int
		quiero int
	}{
		{1, 1}, {2, 0}, {3, 0}, {4, 2}, {5, 10}, {6, 4},
	}

	for _, caso := range casos {
		soluciones := NReinasTodas(caso.n)
		if len(soluciones) != caso.quiero {
			t.Errorf("NReinasTodas(%d) dio %d soluciones, se esperaban %d",
				caso.n, len(soluciones), caso.quiero)
		}
		for _, sol := range soluciones {
			if len(sol) != caso.n || !esSolucionValida(sol) {
				t.Errorf("NReinasTodas(%d) devolvió una solución inválida: %v",
					caso.n, sol)
			}
		}
	}
}

// TestNReinasTodasNoCompartenMemoria comprueba de verdad que las soluciones
// guardadas no comparten el arreglo de respaldo (backing array).
//
// La prueba es por escritura: se guarda una copia de cada solución, después se
// pisa por completo una de ellas y se mira si alguna otra cambió. Si dos
// soluciones apuntaran al mismo arreglo, escribir en una arrastraría a la otra
// y el test falla. Con el código correcto (que copia con make + copy antes de
// guardar) ninguna otra solución se modifica.
//
// Se usa n = 8 a propósito: es el n más chico de los que tienen solución donde
// el problema se manifiesta de forma confiable. Con n = 4, por ejemplo, el
// append del slice parcial reasigna el arreglo y las dos soluciones terminan en
// arreglos distintos aunque falte la copia, así que un test con n = 4 pasaría
// sin detectar nada.
func TestNReinasTodasNoCompartenMemoria(t *testing.T) {
	const n = 8
	soluciones := NReinasTodas(n)

	if len(soluciones) < 2 {
		t.Fatalf("NReinasTodas(%d) dio %d soluciones, se esperaban varias",
			n, len(soluciones))
	}

	if posiciones, hay := indiceDeSolucionesQueCompartenMemoria(soluciones); hay {
		t.Errorf("NReinasTodas(%d): las soluciones %d y %d comparten el mismo "+
			"arreglo; hay que copiar cada solución antes de guardarla",
			n, posiciones[0], posiciones[1])
	}
}

// indiceDeSolucionesQueCompartenMemoria busca dos soluciones que compartan el
// arreglo de respaldo. Devuelve las posiciones de ese par y true si las
// encontró, o un par vacío y false si todas son independientes.
//
// Para descubrirlo pisa por completo cada solución con un valor imposible
// (-1, que no es una columna válida) y compara las demás contra la copia que
// se sacó antes. Después restaura la solución pisada para no ensuciar el
// resultado que devolvió NReinasTodas.
func indiceDeSolucionesQueCompartenMemoria(soluciones [][]int) ([2]int, bool) {
	copias := copiarSoluciones(soluciones)

	for i := range soluciones {
		for j := range soluciones[i] {
			soluciones[i][j] = -1
		}

		for k := range soluciones {
			if k == i {
				continue
			}
			if !iguales(copias[k], soluciones[k]) {
				var posiciones [2]int
				posiciones[0] = i
				posiciones[1] = k
				return posiciones, true
			}
		}

		copiarEnSitio(soluciones[i], copias[i])
	}

	var ninguna [2]int
	return ninguna, false
}

// copiarSoluciones devuelve una copia independiente de cada solución, para
// poder comparar después contra el estado original.
func copiarSoluciones(soluciones [][]int) [][]int {
	copias := make([][]int, len(soluciones))
	for i := range soluciones {
		copias[i] = make([]int, len(soluciones[i]))
		copy(copias[i], soluciones[i])
	}
	return copias
}

// copiarEnSitio copia el contenido de origen dentro de destino, sin cambiar el
// slice destino: se usa para restaurar una solución que se había pisado.
func copiarEnSitio(destino []int, origen []int) {
	for i := range destino {
		destino[i] = origen[i]
	}
}

// iguales indica si dos slices de int tienen el mismo contenido.
func iguales(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// esSolucionValida comprueba que ninguna reina ataque a otra: no comparten
// columna ni ninguna de las dos diagonales.
func esSolucionValida(sol []int) bool {
	for fila := range sol {
		for otra := fila + 1; otra < len(sol); otra++ {
			if sol[fila] == sol[otra] ||
				fila+sol[fila] == otra+sol[otra] ||
				fila-sol[fila] == otra-sol[otra] {
				return false
			}
		}
	}
	return true
}
