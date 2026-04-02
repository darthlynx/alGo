package leetcode3418

import "testing"

func TestMaximumAmount(t *testing.T) {
	tests := []struct {
		coins [][]int
		want  int
	}{
		{
			coins: [][]int{{0, 1, -1}, {1, -2, 3}, {2, -3, 4}},
			want:  8,
		},
		{
			coins: [][]int{{10, 10, 10}, {10, 10, 10}},
			want:  40,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := maximumAmount(tt.coins)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
