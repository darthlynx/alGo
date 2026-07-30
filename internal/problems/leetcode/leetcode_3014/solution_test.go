package leetcode3014

import "testing"

func TestMinimumPushes(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{
		{"single char", "a", 1},
		{"example one round", "abcde", 5},
		{"exactly eight fills first round", "abcdefgh", 8},
		{"nine chars starts second round", "abcdefghi", 10},
		{"example two rounds", "xycdefghij", 12},
		{"all twenty-six letters", "abcdefghijklmnopqrstuvwxyz", 56},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPushes(tt.word); got != tt.want {
				t.Errorf("minimumPushes() = %v, want %v", got, tt.want)
			}
		})
	}
}
