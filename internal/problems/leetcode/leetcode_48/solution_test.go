package leetcode48

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "3x3 example 1",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			name:   "4x4 example 2",
			matrix: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			want:   [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
		{
			name:   "1x1 single element",
			matrix: [][]int{{1}},
			want:   [][]int{{1}},
		},
		{
			name:   "2x2 matrix",
			matrix: [][]int{{1, 2}, {3, 4}},
			want:   [][]int{{3, 1}, {4, 2}},
		},
		{
			name:   "all same values",
			matrix: [][]int{{7, 7, 7}, {7, 7, 7}, {7, 7, 7}},
			want:   [][]int{{7, 7, 7}, {7, 7, 7}, {7, 7, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}
