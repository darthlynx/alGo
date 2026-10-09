package leetcode1541

import (
	"strings"
	"testing"
)

func TestMinInsertions(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "example missing closing parenthesis", s: "(()))", want: 1},
		{name: "example already balanced", s: "())", want: 0},
		{name: "example unmatched parentheses at both ends", s: "))())(", want: 3},
		{name: "empty input", s: "", want: 0},
		{name: "single opening parenthesis", s: "(", want: 2},
		{name: "single closing parenthesis", s: ")", want: 2},
		{name: "opening interrupts closing pair", s: "()())", want: 1},
		{name: "balanced nested and adjacent groups", s: "(())))())", want: 0},
		{name: "odd number of closing parentheses", s: ")))", want: 3},
		{name: "maximum length all opening parentheses", s: strings.Repeat("(", 100000), want: 200000},
		{name: "maximum length all closing parentheses", s: strings.Repeat(")", 100000), want: 50000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}
