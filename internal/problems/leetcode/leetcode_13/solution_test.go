package leetcode13

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example from problem",
			s:    "MCMXCIV",
			want: 1994,
		},
		{
			name: "single numeral",
			s:    "V",
			want: 5,
		},
		{
			name: "subtractive notation",
			s:    "IV",
			want: 4,
		},
		{
			name: "repeated numerals",
			s:    "III",
			want: 3,
		},
		{
			name: "complex case",
			s:    "XLII",
			want: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.s)
			if got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
