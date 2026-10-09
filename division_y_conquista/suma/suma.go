package main

func SumaArray(arreglo []int) int {
	if len(arreglo) == 0 {
		return 0
	}
	if len(arreglo) == 1 {
		return arreglo[0]
	}
	medio := len(arreglo) / 2

	return SumaArray(arreglo[:medio]) + SumaArray(arreglo[medio:])
}
