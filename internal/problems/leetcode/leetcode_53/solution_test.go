package leetcode53

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example from problem",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "all negative numbers",
			nums: []int{-1, -2, -3, -4},
			want: -1,
		},
		{
			name: "single element",
			nums: []int{5},
			want: 5,
		},
		{
			name: "mixed positive and negative",
			nums: []int{2, -1, 2, 3, 4, -5},
			want: 10,
		},
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
