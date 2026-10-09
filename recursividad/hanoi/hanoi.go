package main

// Mover resuelve las torres de Hanoi para n discos: pasa los discos de la
// vara origen a la vara destino usando la vara auxiliar.
//
// No imprime nada: cada movimiento se entrega al callback onMove, que recibe
// el disco movido y las varas de origen y destino. De esa forma la lógica
// queda separada de cómo se muestra el resultado (patrón observador).
func Mover(n int, origen string, aux string, destino string, onMove func(disco int, origen string, destino string)) {
	if n == 1 {
		onMove(1, origen, destino)
		return
	}

	Mover(n-1, origen, destino, aux, onMove)
	onMove(n, origen, destino)
	Mover(n-1, aux, origen, destino, onMove)
}
