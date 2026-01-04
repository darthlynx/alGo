package leetcode1390

import "testing"

func TestSumFourDivisors(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{21, 4, 7},
			want: 32,
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			want: 0,
		},
		{
			nums: []int{21, 21},
			want: 64,
		},
	}

	for _, tc := range tests {
		got := sumFourDivisors(tc.nums)
		if got != tc.want {
			t.Errorf("sumFourDivisors(%v) = %v; want %v", tc.nums, got, tc.want)
		}
	}
}
