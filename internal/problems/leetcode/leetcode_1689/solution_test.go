package leetcode1689

import "testing"

func TestMinPartitions(t *testing.T) {
	tests := []struct {
		name string
		n    string
		want int
	}{
		{"Example 1", "32", 3},
		{"Example 2", "82734", 8},
		{"Example 3", "27346209830709182346", 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPartitions(tt.n); got != tt.want {
				t.Errorf("minPartitions(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
