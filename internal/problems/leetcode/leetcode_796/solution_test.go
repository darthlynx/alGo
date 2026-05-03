package leetcode796

import "testing"

func TestRotateString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		goal string
		want bool
	}{
		{"example 1 - valid rotation", "abcde", "cdeab", true},
		{"example 2 - invalid rotation", "abcde", "abced", false},
		{"same string - zero rotations", "abc", "abc", true},
		{"different lengths", "abc", "abcd", false},
		{"single character match", "a", "a", true},
		{"single character mismatch", "a", "b", false},
		{"empty strings", "", "", true},
		{"all same characters", "aaa", "aaa", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateString(tt.s, tt.goal)
			if got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
