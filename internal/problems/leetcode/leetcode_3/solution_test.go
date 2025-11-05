package leetcode3

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example with repeating characters",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "example with repeating characters 2",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "same character repeated",
			s:    "bbb",
			want: 1,
		},
		{
			name: "empty string",
			s:    "",
			want: 0,
		},
		{
			name: "no repeating characters",
			s:    "abc",
			want: 3,
		},
		{
			name: "two characters",
			s:    "au",
			want: 2,
		},
		{
			name: "complex case 1",
			s:    "dvdf",
			want: 3,
		},
		{
			name: "complex case 2",
			s:    "ckilbkd",
			want: 5,
		},
	}

	solutions := map[string]func(string) int{
		"map solution":    lengthOfLongestSubstringMap,
		"window solution": lengthOfLongestSubstringWindow,
	}

	for solutionName, solution := range solutions {
		t.Run(solutionName, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := solution(tt.s); got != tt.want {
						t.Errorf("%s: lengthOfLongestSubstring() = %v, want %v", solutionName, got, tt.want)
					}
				})
			}
		})
	}
}
