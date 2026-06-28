package leetcode1846

import "testing"

func TestMaximumElementAfterDecrementingAndRearranging(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"example 1", []int{2, 2, 1, 2, 1}, 2},
		{"example 2", []int{100, 1, 1000}, 3},
		{"example 3", []int{1, 2, 3, 4, 5}, 5},
		{"single element", []int{73}, 1},
		{"all ones", []int{1, 1, 1, 1}, 1},
		{"large gaps", []int{1, 1000000}, 2},
		{"already sequential from one", []int{3, 1, 2, 4, 6, 5}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumElementAfterDecrementingAndRearranging(tt.arr); got != tt.want {
				t.Errorf("maximumElementAfterDecrementingAndRearranging() = %v, want %v", got, tt.want)
			}
		})
	}
}
