package leetcode167

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "example from problem",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "target sum at start",
			numbers: []int{0, 0, 3, 4},
			target:  0,
			want:    []int{1, 2},
		},
		{
			name:    "target sum at end",
			numbers: []int{1, 2, 3, 4, 5, 6},
			target:  11,
			want:    []int{5, 6},
		},
		{
			name:    "no solution",
			numbers: []int{1, 2, 3},
			target:  7,
			want:    []int{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.numbers, tt.target)
			if !equalPairs(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// equalPairs checks if two pairs of indices are equal
// order matters since the solution requires first index < second index
func equalPairs(a, b []int) bool {
	if len(a) != 2 || len(b) != 2 {
		return false
	}
	return a[0] == b[0] && a[1] == b[1]
}
