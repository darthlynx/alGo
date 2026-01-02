package leetcode961

import "testing"

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3, 3},
			want: 3,
		},
		{
			nums: []int{2, 1, 2, 5, 3, 2},
			want: 2,
		},
		{
			nums: []int{5, 1, 5, 2, 5, 3, 5, 4},
			want: 5,
		},
		{
			nums: []int{9, 5, 6, 9},
			want: 9,
		},
	}

	for _, tc := range tests {
		got := repeatedNTimes(tc.nums)
		if got != tc.want {
			t.Errorf("Test failed!")
		}
	}
}
