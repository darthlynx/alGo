package leetcode1406

import "testing"

func TestStoneGameIII(t *testing.T) {
	tests := []struct {
		name       string
		stoneValue []int
		want       string
	}{
		{"example 1", []int{1, 2, 3, 7}, "Bob"},
		{"example 2", []int{1, 2, 3, -9}, "Alice"},
		{"example 3", []int{1, 2, 3, 6}, "Tie"},
		{"single positive", []int{5}, "Alice"},
		{"single negative", []int{-5}, "Bob"},
		{"all negatives ties", []int{-1, -2, -3}, "Tie"},
		{"greedy trap", []int{1, 2, 3, -1, -2, -3, 7}, "Alice"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameIII(tt.stoneValue); got != tt.want {
				t.Errorf("stoneGameIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
