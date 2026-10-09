package main

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		houses []int
		want   int
	}{
		{[]int{}, 0},
		{[]int{10}, 10},
		{[]int{2, 10, 3, 6, 8, 1, 7}, 25},
	}
	for _, tt := range tests {
		if got := Rob(tt.houses); got != tt.want {
			t.Errorf("Rob(%v) = %d; want %d", tt.houses, got, tt.want)
		}
	}
}
