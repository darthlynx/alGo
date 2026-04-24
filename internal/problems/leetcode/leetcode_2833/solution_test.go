package leetcode2833

import "testing"

func TestFurthestDistanceFromOrigin(t *testing.T) {
	tests := []struct {
		moves    string
		expected int
	}{
		{"L_RL__R", 3},
		{"_R__LL_", 5},
	}

	for _, test := range tests {
		result := furthestDistanceFromOrigin(test.moves)
		if result != test.expected {
			t.Errorf("For moves %s, expected %d, got %d", test.moves, test.expected, result)
		}
	}
}
