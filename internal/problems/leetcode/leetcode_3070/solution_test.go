package leetcode3070

import "testing"

func TestCountSubmatrices(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want int
	}{
		{"Example 1", [][]int{{7, 6, 3}, {6, 6, 1}}, 18, 4},
		{"Example 2", [][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, 20, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubmatrices(tt.grid, tt.k); got != tt.want {
				t.Errorf("countSubmatrices(%v, %v) = %v, want %v", tt.grid, tt.k, got, tt.want)
			}
		})
	}
}
