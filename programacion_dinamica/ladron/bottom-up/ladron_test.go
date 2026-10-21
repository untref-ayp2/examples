package main

import "testing"

func TestRob(t *testing.T) {
	houses := []int{2, 10, 3, 6, 8, 1, 7}
	if got := Rob(houses); got != 25 {
		t.Errorf("Rob(%v) = %d; want 25", houses, got)
	}
}

func TestRobCalleVacia(t *testing.T) {
	// sin casas no se puede robar nada
	if got := Rob(nil); got != 0 {
		t.Errorf("Rob(nil) = %d; want 0", got)
	}
}

func TestRobUnaCasa(t *testing.T) {
	// con una sola casa, esa es la que se roba
	if got := Rob([]int{5}); got != 5 {
		t.Errorf("Rob([5]) = %d; want 5", got)
	}
}

func TestRobDosCasasConLaPrimeraMayor(t *testing.T) {
	// con dos casas no se pueden robar ambas (son vecinas): gana la primera
	if got := Rob([]int{10, 1}); got != 10 {
		t.Errorf("Rob([10 1]) = %d; want 10", got)
	}
}

func TestRobPrimeraCasaRindeMasQueLaSegunda(t *testing.T) {
	// Mismos casos que en bottom-up-optimizado: las dos versiones aplican la
	// misma decisión, así que tienen que devolver el mismo resultado. Acá el
	// caso sirve para dejar constancia de que la tabla calcula la respuesta
	// correcta cuando la primera casa rinde más que la segunda.
	casos := []struct {
		houses []int
		want   int
	}{
		// Robar la casa 0 y la casa 3: 5 + 5 = 10.
		{[]int{5, 1, 1, 5}, 10},
		// Robar la casa 0 y la casa 3: 100 + 91 = 191.
		{[]int{100, 59, 83, 91}, 191},
	}
	for _, caso := range casos {
		if got := Rob(caso.houses); got != caso.want {
			t.Errorf("Rob(%v) = %d; want %d", caso.houses, got, caso.want)
		}
	}
}
