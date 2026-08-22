package leetcode3622

import "testing"

func TestCheckDivisibility(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"divisible example", 99, true},
		{"not divisible example", 23, false},
		{"single digit", 1, false},
		{"contains zero digit", 10, true},
		{"multi digit not divisible", 12, false},
		{"max constraint", 1000, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDivisibility(tt.n); got != tt.want {
				t.Errorf("checkDivisibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
