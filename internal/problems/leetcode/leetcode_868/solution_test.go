package leetcode868

import "testing"

func TestBinaryGap(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"Example 4", 13, 2},
		{"Example 1", 22, 2},
		{"Example 2", 5, 2},
		{"Example 3", 6, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.n); got != tt.want {
				t.Errorf("binaryGap(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
