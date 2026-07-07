package leetcode3754

import "testing"

func TestSumAndMultiply(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int64
	}{
		{"no zero digits", 123, 738},
		{"single digit", 9, 81},
		{"interior zero", 105, 90},
		{"trailing zero", 10, 1},
		{"all zeros except one", 1000, 1},
		{"repeated digit around zero", 909, 1782},
		{"all nines", 999, 26973},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumAndMultiply(tt.n); got != tt.want {
				t.Errorf("sumAndMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
