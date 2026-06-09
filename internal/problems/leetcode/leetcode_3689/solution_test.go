package leetcode3689

import "testing"

func TestMaxTotalValue(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{
			name: "small array k greater than 1",
			nums: []int{1, 3, 2},
			k:    2,
			want: 4,
		},
		{
			name: "single element",
			nums: []int{5},
			k:    10,
			want: 0,
		},
		{
			name: "all same values",
			nums: []int{7, 7, 7, 7},
			k:    5,
			want: 0,
		},
		{
			name: "mixed with negatives",
			nums: []int{-3, 5, 2},
			k:    3,
			want: 24,
		},
		{
			name: "k equals one",
			nums: []int{1, 5, 3, 9, 2},
			k:    1,
			want: 8,
		},
		{
			name: "large spread and k",
			nums: []int{1, 5, 3, 9, 2},
			k:    4,
			want: 32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTotalValue(tt.nums, tt.k); got != tt.want {
				t.Errorf("maxTotalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
