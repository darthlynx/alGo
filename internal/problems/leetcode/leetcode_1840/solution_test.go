package leetcode1840

import "testing"

func TestMaxBuilding(t *testing.T) {
	tests := []struct {
		name         string
		n            int
		restrictions [][]int
		want         int
	}{
		{
			name:         "example 1",
			n:            5,
			restrictions: [][]int{{2, 1}, {4, 1}},
			want:         2,
		},
		{
			name:         "example 2 no restrictions",
			n:            6,
			restrictions: [][]int{},
			want:         5,
		},
		{
			name:         "example 3 multiple restrictions",
			n:            10,
			restrictions: [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}},
			want:         5,
		},
		{
			name:         "single building",
			n:            1,
			restrictions: [][]int{},
			want:         0,
		},
		{
			name:         "restriction at last building",
			n:            5,
			restrictions: [][]int{{5, 0}},
			want:         2,
		},
		{
			name:         "very high restriction does not bind",
			n:            10,
			restrictions: [][]int{{3, 1000000000}},
			want:         9,
		},
		{
			name:         "unsorted dense restrictions",
			n:            10,
			restrictions: [][]int{{8, 5}, {9, 0}, {6, 2}, {4, 0}, {3, 2}, {10, 0}, {5, 3}, {7, 3}, {2, 4}},
			want:         2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxBuilding(tt.n, tt.restrictions); got != tt.want {
				t.Errorf("maxBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}
