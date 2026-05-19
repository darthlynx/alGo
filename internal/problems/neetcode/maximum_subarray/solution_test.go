package maximumsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example from problem", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"single element", []int{1}, 1},
		{"all positive", []int{5, 4, -1, 7, 8}, 23},
		{"all negative", []int{-3, -2, -1}, -1},
		{"single negative", []int{-1}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubArray(tt.nums)
			if got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
