package leetcode2540

import "testing"

func TestGetCommon(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{"example 1 - common at start", []int{1, 2, 3}, []int{2, 4}, 2},
		{"example 2 - common at end", []int{1, 2, 3, 6}, []int{2, 3, 4, 5}, 2},
		{"no common value", []int{1, 3, 5}, []int{2, 4, 6}, -1},
		{"single element match", []int{7}, []int{7}, 7},
		{"single element no match", []int{1}, []int{2}, -1},
		{"common is the smallest", []int{1, 4, 7}, []int{1, 2, 3}, 1},
		{"common is the largest", []int{2, 5, 8}, []int{3, 6, 8}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCommon(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("getCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}
