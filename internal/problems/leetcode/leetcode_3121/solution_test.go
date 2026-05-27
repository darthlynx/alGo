package leetcode3121

import "testing"

func TestNumberOfSpecialChars(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{
		{"example 1 - only a is special", "aaAbBb", 1},
		{"example 2 - no uppercase", "abc", 0},
		{"example 3 - uppercase before lowercase", "AbBCab", 0},
		{"both letters special", "aAbB", 2},
		{"uppercase before lowercase not special", "Aa", 0},
		{"empty string", "", 0},
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
