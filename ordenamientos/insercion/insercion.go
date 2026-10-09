package main

// OrdenarInsercion ordena el slice en orden ascendente con el algoritmo de
// inserción: construye el slice ordenado un elemento a la vez, insertando
// cada nuevo elemento en su posición correcta entre los ya ordenados, como
// se ordenan cartas en la mano. Es estable, in-place y online.
// Complejidad: O(n²) peor caso, O(n) mejor caso.
func OrdenarInsercion(arr []int) {
	for i := 1; i < len(arr); i++ {
		clave := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > clave {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = clave
	}
}
