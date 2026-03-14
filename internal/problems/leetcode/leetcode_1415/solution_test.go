package leetcode1415

import "testing"

func TestGetHappyString(t *testing.T) {
	tests := []struct {
		n        int
		k        int
		expected string
	}{
		{n: 1, k: 3, expected: "c"},
		{n: 1, k: 4, expected: ""},
		{n: 3, k: 9, expected: "cab"},
	}

	for _, test := range tests {
		if result := getHappyString(test.n, test.k); result != test.expected {
			t.Errorf("getHappyString(%d, %d) = %s; expected %s", test.n, test.k, result, test.expected)
		}
	}
}
