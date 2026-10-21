package main

// OrdenarInsercion ordena el slice en orden ascendente con el algoritmo de
// inserción: construye el slice ordenado un elemento a la vez, insertando
// cada nuevo elemento en su posición correcta entre los ya ordenados, como
// se ordenan cartas en la mano. Es estable, in-place y online.
// Complejidad: O(n²) peor caso, O(n) mejor caso.
func OrdenarInsercion(arr []int) {
	// arr[0..i-1] ya está ordenado; se inserta arr[i] en su lugar.
	for i := 1; i < len(arr); i++ {
		clave := arr[i]
		j := i - 1
		// Se corren a la derecha todos los mayores que la clave. Como el
		// corrimiento se hace de atrás hacia adelante, no se pisa ningún
		// valor que todavía falte leer.
		for j >= 0 && arr[j] > clave {
			arr[j+1] = arr[j]
			j--
		}
		// j quedó en la última posición cuya clave no es mayor: la clave
		// va justo después, en j+1 (que puede ser la posición original i,
		// si no hubo que correr nada).
		arr[j+1] = clave
	}
}
