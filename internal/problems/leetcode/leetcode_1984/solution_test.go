package leetcode1984

import (
	"fmt"
	"testing"
)

func TestMinimumDifference(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{90},
			k:    1,
			want: 0,
		},
		{
			nums: []int{9, 4, 1, 7},
			k:    2,
			want: 2,
		},
		{
			nums: []int{10, 100, 300, 200, 1000, 20, 30},
			k:    3,
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums=%v,k=%d", tt.nums, tt.k), func(t *testing.T) {
			got := minimumDifference(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
