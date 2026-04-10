package leetcode3740

import "testing"

func TestMinimumDistance(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1,2,1,1,3}, 6},
		{[]int{1,1,2,3,2,1,2}, 8},
		{[]int{1}, -1},
	}

	for _, test := range tests {
		if result := minimumDistance(test.nums); result != test.expected {
			t.Errorf("For %v expected %d but got %d", test.nums, test.expected, result)
		}
	}
}
