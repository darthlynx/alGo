package leetcode3047

import "testing"

func TestLargestSquareArea(t *testing.T) {
	tests := []struct {
		bottomLeft [][]int
		topRight   [][]int
		expected   int64
	}{
		{
			bottomLeft: [][]int{{0, 0}, {2, 2}, {3, 1}},
			topRight:   [][]int{{3, 3}, {4, 4}, {6, 6}},
			expected:   1,
		},
		{
			bottomLeft: [][]int{{0, 0}, {0, 0}},
			topRight:   [][]int{{1, 1}, {1, 1}},
			expected:   1,
		},
		{
			bottomLeft: [][]int{{0, 0}, {2, 1}},
			topRight:   [][]int{{1, 2}, {3, 3}},
			expected:   0,
		},
		{
			bottomLeft: [][]int{{2, 4}, {1, 1}},
			topRight:   [][]int{{4, 5}, {3, 2}},
			expected:   0,
		},
	}

	for _, test := range tests {
		actual := largestSquareArea(test.bottomLeft, test.topRight)
		if actual != test.expected {
			t.Errorf("For bottomLeft=%v and topRight=%v, expected %d but got %d",
				test.bottomLeft, test.topRight, test.expected, actual)
		}
	}
}
