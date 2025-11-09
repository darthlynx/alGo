package leetcode121

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "example from problem",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "no profit",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "single day",
			prices: []int{5},
			want:   0,
		},
		{
			name:   "profit at the end",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
