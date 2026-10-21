package main

import "testing"

func TestRob(t *testing.T) {
	houses := []int{2, 10, 3, 6, 8, 1, 7}
	if got := Rob(houses); got != 25 {
		t.Errorf("Rob(%v) = %d; want 25", houses, got)
	}
}

func TestRobPrimeraCasaRindeMasQueLaSegunda(t *testing.T) {
	// Casos de regresión que distinguen la inicialización correcta de
	// `prev` (= table[1] = max(houses[0], houses[1])) de la incorrecta
	// (= houses[1]). Si `prev` arranca con houses[1] y la primera casa
	// rinde más que la segunda, la versión optimizada arrastra un valor
	// menor y devuelve de menos. Estos arreglos tienen al menos tres
	// casas, así que llegan al bucle.
	casos := []struct {
		houses []int
		want   int
	}{
		// La mejor jugada es robar la casa 0 y la casa 3: 5 + 5 = 10.
		// Con la inicialización incorrecta se devuelve 6.
		{[]int{5, 1, 1, 5}, 10},
		// Robando 100 + 91 = 191; con la incorrecta se devuelve 183.
		{[]int{100, 59, 83, 91}, 191},
	}
	for _, caso := range casos {
		if got := Rob(caso.houses); got != caso.want {
			t.Errorf("Rob(%v) = %d; want %d", caso.houses, got, caso.want)
		}
	}
}

func TestRobCasasPocas(t *testing.T) {
	casos := []struct {
		houses []int
		want   int
	}{
		{[]int{}, 0},
		{[]int{2}, 2},
		{[]int{2, 10}, 10},
		{[]int{10, 2}, 10},
	}
	for _, caso := range casos {
		if got := Rob(caso.houses); got != caso.want {
			t.Errorf("Rob(%v) = %d; want %d", caso.houses, got, caso.want)
		}
	}
}
