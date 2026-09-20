package leetcode_3498

import "testing"

func TestReverseDegree(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"example abc", "abc", 148},
		{"example zaz", "zaz", 56},
		{"single a", "a", 26},
		{"single z", "z", 1},
		{"all same", "aa", 78},
		{"descending values", "zyx", 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseDegree(tt.s); got != tt.want {
				t.Errorf("reverseDegree() = %v, want %v", got, tt.want)
			}
		})
	}
}
