package leetcode678

import (
	"strings"
	"testing"
)

func TestCheckValidString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "example valid pair", s: "()", want: true},
		{name: "example wildcard as empty", s: "(*)", want: true},
		{name: "example wildcard as opening parenthesis", s: "(*))", want: true},
		{name: "example single opening parenthesis", s: "(", want: false},
		{name: "empty input", s: "", want: true},
		{name: "single closing parenthesis", s: ")", want: false},
		{name: "single wildcard", s: "*", want: true},
		{name: "wildcard as closing parenthesis", s: "(*", want: true},
		{name: "closing parenthesis before wildcard", s: ")*", want: false},
		{name: "opening parenthesis after wildcard", s: "*(", want: false},
		{name: "balanced counts in invalid order", s: ")(", want: false},
		{name: "wildcards serving different roles", s: "(*))((*)", want: true},
		{name: "too few wildcards for opening parentheses", s: "(((**", want: false},
		{name: "too few wildcards for closing parentheses", s: "**)))", want: false},
		{name: "maximum length all wildcards", s: strings.Repeat("*", 100), want: true},
		{name: "maximum length all opening parentheses", s: strings.Repeat("(", 100), want: false},
		{name: "maximum length all closing parentheses", s: strings.Repeat(")", 100), want: false},
		{name: "maximum length nested parentheses", s: strings.Repeat("(", 50) + strings.Repeat(")", 50), want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
