package main

// OrdenarSeleccion ordena el slice en orden ascendente con el algoritmo de
// selección: encuentra repetidamente el elemento mínimo del resto no ordenado
// y lo coloca al principio de la lista. Es in-place y no es estable.
// Complejidad: O(n²) en todos los casos.
func OrdenarSeleccion(arr []int) {
	// El bucle llega hasta len(arr)-1 porque cuando queda un solo elemento
	// sin ordenar, ya está en su lugar.
	for i := 0; i < len(arr)-1; i++ {
		// Buscar el mínimo del tramo que todavía no está ordenado.
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Intercambiar el mínimo con la posición i; si ya está en su lugar
		// el intercambio no cambia nada, así que se evita (como en la
		// versión visual).
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}
