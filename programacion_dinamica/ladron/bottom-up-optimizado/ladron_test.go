package main

import "testing"

func TestRob(t *testing.T) {
	houses := []int{2, 10, 3, 6, 8, 1, 7}
	if got := Rob(houses); got != 25 {
		t.Errorf("Rob(%v) = %d; want 25", houses, got)
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
