package leetcode3612

import "testing"

func TestProcessStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"only regular chars", "abc", "abc"},
		{"star removes last char", "ab*", "a"},
		{"star on empty stack is no-op", "*a", "a"},
		{"hash duplicates current result", "ab#", "abab"},
		{"percent reverses current result", "abc%", "cba"},
		{"combined operations", "ab#c%", "cbaba"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processStr(tt.s)
			if got != tt.want {
				t.Errorf("processStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
