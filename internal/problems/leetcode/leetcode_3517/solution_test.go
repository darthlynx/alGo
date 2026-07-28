package leetcode3517

import "testing"

func TestSmallestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"single char", "z", "z"},
		{"odd length example", "babab", "abbba"},
		{"even length example", "daccad", "acddca"},
		{"all same", "cccc", "cccc"},
		{"already smallest", "abcba", "abcba"},
		{"reorders to smaller", "adbbda", "abddba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestPalindrome(tt.s); got != tt.want {
				t.Errorf("smallestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
