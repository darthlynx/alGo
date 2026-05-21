package leetcode3043

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		want int
	}{
		{"example 1 - prefix chain", []int{1, 10, 100}, []int{1000}, 3},
		{"example 2 - no common prefix", []int{1, 2, 3}, []int{4, 4, 4}, 0},
		{"full match single digit", []int{5}, []int{5}, 1},
		{"prefix shorter in arr2", []int{123, 456}, []int{12, 45}, 2},
		{"multiple matches take longest", []int{12, 123, 1234}, []int{12345}, 4},
		{"single element each no match", []int{1}, []int{2}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.arr1, tt.arr2)
			if got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
