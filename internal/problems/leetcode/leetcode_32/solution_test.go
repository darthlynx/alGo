package leetcode_32

import (
	"strings"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example with unmatched opening parenthesis",
			s:    "(()",
			want: 2,
		},
		{
			name: "example with unmatched closing parentheses",
			s:    ")()())",
			want: 4,
		},
		{
			name: "empty input",
			s:    "",
			want: 0,
		},
		{
			name: "single opening parenthesis",
			s:    "(",
			want: 0,
		},
		{
			name: "single closing parenthesis",
			s:    ")",
			want: 0,
		},
		{
			name: "single valid pair",
			s:    "()",
			want: 2,
		},
		{
			name: "balanced counts in invalid order",
			s:    ")(",
			want: 0,
		},
		{
			name: "adjacent and nested valid groups",
			s:    "()(())",
			want: 6,
		},
		{
			name: "valid groups separated by unmatched closing parenthesis",
			s:    "())(())",
			want: 4,
		},
		{
			name: "valid groups separated by unmatched opening parenthesis",
			s:    "(())(()",
			want: 4,
		},
		{
			name: "valid substring between invalid ends",
			s:    ")(()())(",
			want: 6,
		},
		{
			name: "maximum length all opening parentheses",
			s:    strings.Repeat("(", 30000),
			want: 0,
		},
		{
			name: "maximum length all closing parentheses",
			s:    strings.Repeat(")", 30000),
			want: 0,
		},
		{
			name: "maximum length nested valid parentheses",
			s:    strings.Repeat("(", 15000) + strings.Repeat(")", 15000),
			want: 30000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
