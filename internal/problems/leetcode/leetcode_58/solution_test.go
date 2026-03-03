package leetcode58

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"Example 1", "Hello World", 5},
		{"Example 2", "   fly me   to   the moon  ", 4},
		{"Example 3", "luffy is still joyboy", 6},
		{"Example 4", "a ", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("lengthOfLastWord(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
