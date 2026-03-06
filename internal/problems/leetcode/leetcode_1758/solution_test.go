package leetcode1758

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"Example 1", "0100", 1},
		{"Example 2", "10", 0},
		{"Example 3", "1111", 2},
		{"Example 4", "10010100", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.s); got != tt.want {
				t.Errorf("minOperations(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
