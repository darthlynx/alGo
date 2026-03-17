package leetcode1727

import "testing"

func TestLargestSubmatrix(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected int
	}{
		{matrix: [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}, expected: 4},
		{matrix: [][]int{{1, 0, 1, 0, 1}}, expected: 3},
	}

	for _, test := range tests {
		if result := largestSubmatrix(test.matrix); result != test.expected {
			t.Errorf("largestSubmatrix(%v) = %d; expected %d", test.matrix, result, test.expected)
		}
	}
}
