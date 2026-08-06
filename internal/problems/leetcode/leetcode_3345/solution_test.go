package leetcode3345

import "testing"

func TestSmallestNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		t    int
		want int
	}{
		{"example 1", 10, 2, 10},
		{"example 2", 15, 3, 16},
		{"t is one returns n", 7, 1, 7},
		{"needs increment for prime factor", 1, 5, 5},
		{"t is ten needs zero digit", 1, 10, 10},
		{"already divisible no increment", 25, 10, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNumber(tt.n, tt.t); got != tt.want {
				t.Errorf("smallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
