package leetcode1855

import "testing"

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{"Example 1", []int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}, 2},
		{"Example 2", []int{2, 2, 2}, []int{10, 10, 1}, 1},
		{"Example 3", []int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
