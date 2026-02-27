package leetcode1404

import "testing"

func TestNumSteps(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"Example 1", "1101", 6},
		{"Example 2", "10", 1},
		{"Example 3", "1", 0},
		{"Example 4", "11", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSteps(tt.s); got != tt.want {
				t.Errorf("numSteps(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
