package main

// Mover resuelve las torres de Hanoi para n discos: pasa los discos de la
// vara origen a la vara destino usando la vara auxiliar.
//
// No imprime nada: cada movimiento se entrega al callback onMove, que recibe
// el disco movido y las varas de origen y destino. De esa forma la lógica
// queda separada de cómo se muestra el resultado (patrón observador).
//
// Recursión: para mover n discos, primero mueve los n-1 de arriba a la
// auxiliar, después mueve el disco n y por último pasa los n-1 de la
// auxiliar al destino.
//
// Caso base: n == 1 mueve el disco directo. Termina porque cada llamada
// recursiva reduce n en 1: hay n niveles de recursión y dos llamadas por
// nivel (2^n - 1 movimientos en total).
//
// Precondición: n >= 1.
func Mover(n int, origen string, aux string, destino string, onMove func(disco int, origen string, destino string)) {
	if n == 1 {
		onMove(1, origen, destino)
		return
	}

	Mover(n-1, origen, destino, aux, onMove)
	onMove(n, origen, destino)
	Mover(n-1, aux, origen, destino, onMove)
}
