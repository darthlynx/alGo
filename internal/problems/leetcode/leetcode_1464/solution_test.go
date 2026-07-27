package leetcode1464

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{3, 4, 5, 2}, 12},
		{"example 2", []int{1, 5, 4, 5}, 16},
		{"two elements", []int{3, 7}, 12},
		{"all same", []int{4, 4, 4, 4}, 9},
		{"all ones", []int{1, 1, 1}, 0},
		{"two largest not adjacent in input", []int{10, 2, 9, 3}, 72},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
