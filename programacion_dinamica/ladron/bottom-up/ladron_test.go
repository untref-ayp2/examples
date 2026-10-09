package main

import "testing"

func TestRob(t *testing.T) {
	houses := []int{2, 10, 3, 6, 8, 1, 7}
	if got := Rob(houses); got != 25 {
		t.Errorf("Rob(%v) = %d; want 25", houses, got)
	}
}
