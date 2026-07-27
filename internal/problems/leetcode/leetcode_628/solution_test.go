package leetcode628

import "testing"

func TestMaximumProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{1, 2, 3}, 6},
		{"example 2", []int{1, 2, 3, 4}, 24},
		{"all negatives", []int{-1, -2, -3}, -6},
		{"two negatives beat three positives", []int{-4, -3, 1, 2, 3}, 36},
		{"negatives with zero", []int{-10, -10, 0, 1}, 100},
		{"five negatives", []int{-5, -4, -3, -2, -1}, -6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
