package leetcode1536

import "testing"

func TestMinSwaps(t *testing.T) {
	tests := []struct {
		name string
		nums [][]int
		want int
	}{
		{"Example 1", [][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}, 3},
		{"Example 2", [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwaps(tt.nums); got != tt.want {
				t.Errorf("minSwaps(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
