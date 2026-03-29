package leetcode2839

import "testing"

func TestCanBeEqual(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{
			s1:   "abcd",
			s2:   "cdab",
			want: true,
		},
		{
			s1:   "abcd",
			s2:   "dacb",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := canBeEqual(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanBeEqual2(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{s1: "abcd", s2: "cdab", want: true},
		{s1: "abcd", s2: "abcd", want: true},
		{s1: "abcd", s2: "dacb", want: false},
		{s1: "abcd", s2: "abdc", want: false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := canBeEqual_2(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("canBeEqual_2(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
