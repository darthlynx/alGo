package leetcode1260

import (
	"reflect"
	"testing"
)

func TestShiftGrid(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want [][]int
	}{
		{
			name: "example 1",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:    1,
			want: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			name: "example 2",
			grid: [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
			k:    4,
			want: [][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},
		{
			name: "k equals total wraps to original",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:    9,
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			name: "k zero no shift",
			grid: [][]int{{1, 2}, {3, 4}},
			k:    0,
			want: [][]int{{1, 2}, {3, 4}},
		},
		{
			name: "k larger than total",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:    10,
			want: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			name: "single cell",
			grid: [][]int{{7}},
			k:    5,
			want: [][]int{{7}},
		},
		{
			name: "single row",
			grid: [][]int{{1, 2, 3, 4}},
			k:    2,
			want: [][]int{{3, 4, 1, 2}},
		},
		{
			name: "non-square matrix",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}},
			k:    2,
			want: [][]int{{5, 6, 1}, {2, 3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftGrid(tt.grid, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shiftGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
