package main

// Este archivo resuelve las torres de Hanoi sin animación y sin callbacks:
// la función devuelve la lista completa de movimientos.

// Movimiento describe el traslado de un disco de una vara a otra. Disco es el
// tamaño del disco movido (1 es el más chico) y Origen y Destino son los
// nombres de las varas ("A", "B" o "C").
type Movimiento struct {
	Disco   int
	Origen  string
	Destino string
}

// Mover resuelve las torres de Hanoi para n discos: devuelve, en orden, los
// movimientos necesarios para pasar la pila de la vara origen a la vara
// destino usando la vara auxiliar.
//
// Precondición: n >= 1 y los tres nombres de vara son distintos entre sí
// (convención: "A", "B" y "C").
//
// Devuelve 2^n - 1 movimientos: cada disco duplica el trabajo de los n-1 de
// arriba más su propio movimiento.
func Mover(n int, origen string, aux string, destino string) []Movimiento {
	var movimientos []Movimiento
	mover(n, origen, aux, destino, &movimientos)
	return movimientos
}

// mover es la recursión: para pasar n discos de origen a destino, primero
// pasa los n-1 de arriba a la auxiliar, después mueve el disco n y por último
// pasa los n-1 de la auxiliar al destino.
//
// Acumula los movimientos en la lista apuntada por movimientos (por eso el
// puntero: es el equivalente sin closures de "ir colectando lo que pasa").
//
// Termina porque cada llamada recursiva reduce n en 1 y el caso base es
// n == 1: hay exactamente n niveles de recursión y dos llamadas por nivel.
func mover(n int, origen string, aux string, destino string, movimientos *[]Movimiento) {
	// Caso base: con un solo disco, se mueve directo de origen a destino.
	if n == 1 {
		*movimientos = append(*movimientos, Movimiento{Disco: n, Origen: origen, Destino: destino})
		return
	}

	// Paso 1: los n-1 discos de arriba van a la auxiliar (el destino hace
	// de auxiliar en este paso).
	mover(n-1, origen, destino, aux, movimientos)

	// Paso 2: el disco más grande va de origen a destino.
	*movimientos = append(*movimientos, Movimiento{Disco: n, Origen: origen, Destino: destino})

	// Paso 3: los n-1 discos que están en la auxiliar van al destino (ahora
	// el origen hace de auxiliar).
	mover(n-1, aux, origen, destino, movimientos)
}
