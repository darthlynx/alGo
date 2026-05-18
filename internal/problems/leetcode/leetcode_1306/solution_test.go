package leetcode1306

import "testing"

func TestCanReach(t *testing.T) {
	tests := []struct {
		name  string
		arr   []int
		start int
		want  bool
	}{
		{
			name:  "example 1: reachable via multiple jumps",
			arr:   []int{4, 2, 3, 0, 3, 1, 2},
			start: 5,
			want:  true,
		},
		{
			name:  "example 2: start far from zero but reachable",
			arr:   []int{4, 2, 3, 0, 3, 1, 2},
			start: 0,
			want:  true,
		},
		{
			name:  "example 3: no path to a zero",
			arr:   []int{3, 0, 2, 1, 2},
			start: 2,
			want:  false,
		},
		{
			name:  "start index already zero",
			arr:   []int{0},
			start: 0,
			want:  true,
		},
		{
			name:  "single element non-zero with no reachable zero",
			arr:   []int{1},
			start: 0,
			want:  false,
		},
		{
			name:  "zero reachable only via left jump",
			arr:   []int{0, 1, 2},
			start: 2,
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrCopy := make([]int, len(tt.arr))
			copy(arrCopy, tt.arr)
			got := canReach(arrCopy, tt.start)
			if got != tt.want {
				t.Errorf("canReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
