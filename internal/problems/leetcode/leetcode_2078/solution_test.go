package leetcode2078

import "testing"

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		name   string
		colors []int
		want   int
	}{
		{"Example 1", []int{1, 1, 1, 6, 1, 1, 1}, 3},
		{"Example 2", []int{1, 8, 3, 8, 3}, 4},
		{"Example 3", []int{0, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.colors); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
