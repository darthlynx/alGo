package leetcode3090

import "testing"

func TestMaximumLengthSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"example 1", "bcbbbcba", 4},
		{"example 2", "aaaa", 2},
		{"single char", "a", 1},
		{"all distinct", "abcdef", 6},
		{"three same in a row", "aaa", 2},
		{"pair then triple", "aabbb", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLengthSubstring(tt.s); got != tt.want {
				t.Errorf("maximumLengthSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
