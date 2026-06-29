package leetcode1967

import "testing"

func TestNumOfStrings(t *testing.T) {
	tests := []struct {
		name     string
		patterns []string
		word     string
		want     int
	}{
		{"example 1", []string{"a", "abc", "bc", "d"}, "abc", 3},
		{"example 2", []string{"a", "b", "c"}, "aaaaabbbbb", 2},
		{"example 3", []string{"a", "a", "a"}, "ab", 3},
		{"none match", []string{"x", "yz"}, "abc", 0},
		{"all match", []string{"ab", "bc", "abc"}, "abc", 3},
		{"single pattern", []string{"abc"}, "abc", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfStrings(tt.patterns, tt.word); got != tt.want {
				t.Errorf("numOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
