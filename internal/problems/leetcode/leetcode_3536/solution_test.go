package leetcode3536

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"two digits", 31, 3},
		{"repeated digit", 22, 4},
		{"three digits", 124, 8},
		{"all same digits", 9999, 81},
		{"contains zero", 105, 5},
		{"largest digits at ends", 987654321, 72},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.n); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
