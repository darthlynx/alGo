package leetcode1975

import "testing"

func TestMaxMatrixSum(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int64
	}{
		{
			matrix: [][]int{
				{1, -1},
				{-1, 1},
			},
			want: 4,
		},
		{
			matrix: [][]int{
				{1, 2, 3},
				{-1, -2, -3},
				{1, 2, 3},
			},
			want: 16,
		},
		{
			matrix: [][]int{
				{-1, -1, -1},
				{-1, 2, 3},
				{-1, -2, -3},
			},
			want: 13,
		},
	}

	for _, tc := range tests {
		got := maxMatrixSum(tc.matrix)
		if got != tc.want {
			t.Errorf("wanted %d, got %d", tc.want, got)
		}
	}
}
