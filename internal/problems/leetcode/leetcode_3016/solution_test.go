package leetcode3016

import "testing"

func TestMinimumPushes(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{
		{"example 1", "abcde", 5},
		{"example 2", "xyzxyzxyzxyz", 12},
		{"example 3", "aabbccddeeffgghhiiiiii", 24},
		{"single char", "a", 1},
		{"all same char", "aaaa", 4},
		{"exactly eight distinct", "abcdefgh", 8},
		{"nine distinct spills to second push", "abcdefghi", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPushes(tt.word); got != tt.want {
				t.Errorf("minimumPushes() = %v, want %v", got, tt.want)
			}
		})
	}
}
