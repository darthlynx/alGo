package leetcode2033

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		x    int
		want int
	}{
		{
			name: "example 1 - evenly spaced 2x2",
			grid: [][]int{{2, 4}, {6, 8}},
			x:    2,
			want: 4,
		},
		{
			name: "example 2 - impossible different remainders",
			grid: [][]int{{1, 5}, {2, 8}},
			x:    2,
			want: -1,
		},
		{
			name: "example 3 - x equals one",
			grid: [][]int{{1, 2}, {3, 4}},
			x:    1,
			want: 4,
		},
		{
			name: "single element",
			grid: [][]int{{5}},
			x:    3,
			want: 0,
		},
		{
			name: "all same values",
			grid: [][]int{{3, 3}, {3, 3}},
			x:    2,
			want: 0,
		},
		{
			name: "one element with incompatible remainder",
			grid: [][]int{{2, 4}, {6, 9}},
			x:    2,
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations(tt.grid, tt.x)
			if got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
