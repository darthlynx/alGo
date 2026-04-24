package maximumsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	for _, test := range tests {
		result := maxSubArray(test.nums)
		if result != test.expected {
			t.Errorf("For nums %v, expected %d, got %d", test.nums, test.expected, result)
		}
	}
}
