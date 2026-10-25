package bfs

import "testing"

func TestClosedIsland(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example 1 multiple closed and open islands",
			grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0},
			},
			want: 2,
		},
		{
			name: "example 2 diagonal land does not connect islands",
			grid: [][]int{
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0},
			},
			want: 1,
		},
		{
			name: "example 3 nested closed islands",
			grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1},
			},
			want: 2,
		},
		{
			name: "single land cell touches boundary",
			grid: [][]int{{0}},
			want: 0,
		},
		{
			name: "single water cell",
			grid: [][]int{{1}},
			want: 0,
		},
		{
			name: "all land is one open island",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 0,
		},
		{
			name: "all water has no islands",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			want: 0,
		},
		{
			name: "open island branches are all visited",
			grid: [][]int{
				{1, 0, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 1, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closedIsland(tt.grid); got != tt.want {
				t.Errorf("closedIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
