package leetcode1266

import "testing"

func TestMinTimeToVisitAllPoints(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "test case 1",
			points: [][]int{{1, 1}, {3, 4}, {-1, 0}},
			want:   7,
		},
		{
			name:   "test case 2",
			points: [][]int{{3, 2}, {-2, 2}},
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minTimeToVisitAllPoints(tt.points)
			if got != tt.want {
				t.Errorf("minTimeToVisitAllPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
