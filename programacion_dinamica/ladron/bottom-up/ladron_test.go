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
