package leetcode921

import (
	"strings"
	"testing"
)

func TestMinAddToMakeValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "example extra closing parenthesis", s: "())", want: 1},
		{name: "example all opening parentheses", s: "(((", want: 3},
		{name: "empty input", s: "", want: 0},
		{name: "single opening parenthesis", s: "(", want: 1},
		{name: "single closing parenthesis", s: ")", want: 1},
		{name: "balanced nested and adjacent pairs", s: "(())()", want: 0},
		{name: "closing before opening", s: ")(", want: 2},
		{name: "unmatched parentheses at both ends", s: "))(())((", want: 4},
		{name: "maximum length all closing parentheses", s: strings.Repeat(")", 1000), want: 1000},
		{name: "maximum length all opening parentheses", s: strings.Repeat("(", 1000), want: 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAddToMakeValid(tt.s); got != tt.want {
				t.Errorf("minAddToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
