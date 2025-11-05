package leetcode1

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example from problem",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{2, 1},
		},
		{
			name:   "target at start",
			nums:   []int{1, 2, 3, 4},
			target: 3,
			want:   []int{1, 0},
		},
		{
			name:   "no solution",
			nums:   []int{1, 2, 3},
			target: 7,
			want:   []int{-1, -1},
		},
		{
			name:   "same numbers",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !equalPairs(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// equalPairs checks if two pairs of indices represent the same solution
// order doesn't matter since there's exactly one solution
func equalPairs(a, b []int) bool {
	if len(a) != 2 || len(b) != 2 {
		return false
	}
	return (a[0] == b[0] && a[1] == b[1]) || (a[0] == b[1] && a[1] == b[0])
}
