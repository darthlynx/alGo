package leetcode3120

import "testing"

func TestNumberOfSpecialChars(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{
		{"multiple special chars", "aaAbcBC", 3},
		{"no uppercase", "abc", 0},
		{"repeated pair reappears later", "abBCab", 1},
		{"empty string", "", 0},
		{"single pair", "aA", 1},
		{"repeated chars both cases", "aaAA", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfSpecialChars(tt.word)
			if got != tt.want {
				t.Errorf("numberOfSpecialChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
