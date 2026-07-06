package leetcode1288

import "testing"

func TestRemoveCoveredIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "example 1",
			intervals: [][]int{{1, 4}, {3, 6}, {2, 8}},
			want:      2,
		},
		{
			name:      "example 2",
			intervals: [][]int{{1, 4}, {2, 3}},
			want:      1,
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 10}},
			want:      1,
		},
		{
			name:      "no coverage",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:      3,
		},
		{
			name:      "same left boundary keeps widest",
			intervals: [][]int{{1, 2}, {1, 4}, {1, 3}},
			want:      1,
		},
		{
			name:      "all covered by first",
			intervals: [][]int{{1, 100}, {2, 3}, {4, 5}, {6, 7}},
			want:      1,
		},
		{
			name:      "chained overlaps not covered",
			intervals: [][]int{{1, 4}, {3, 6}, {5, 8}},
			want:      3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeCoveredIntervals(tt.intervals); got != tt.want {
				t.Errorf("removeCoveredIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
