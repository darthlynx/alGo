package leetcode9

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"Example 1", 121, true},
		{"Example 2", -121, false},
		{"Example 3", 10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome(%v) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
