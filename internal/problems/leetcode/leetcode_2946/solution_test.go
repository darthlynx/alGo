package leetcode2946

import "testing"

func TestAreSimilar(t *testing.T) {
	tests := []struct {
		mat      [][]int
		k        int
		expected bool
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 4, false},
		{[][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}}, 2, true},
	}

	for i, test := range tests {
		if result := areSimilar(test.mat, test.k); result != test.expected {
			t.Errorf("Test %d: expected %t, got %t", i+1, test.expected, result)
		}
	}
}
