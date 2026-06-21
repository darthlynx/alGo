package leetcode1833

import "testing"

func TestMaxIceCream(t *testing.T) {
	tests := []struct {
		name  string
		costs []int
		coins int
		want  int
	}{
		{
			name:  "example 1",
			costs: []int{1, 3, 2, 4, 1},
			coins: 7,
			want:  4,
		},
		{
			name:  "example 2 cannot afford any",
			costs: []int{10, 6, 8, 7, 7, 8},
			coins: 5,
			want:  0,
		},
		{
			name:  "example 3 buy all",
			costs: []int{1, 6, 3, 1, 2, 5},
			coins: 20,
			want:  6,
		},
		{
			name:  "single bar affordable",
			costs: []int{5},
			coins: 5,
			want:  1,
		},
		{
			name:  "single bar too expensive",
			costs: []int{5},
			coins: 4,
			want:  0,
		},
		{
			name:  "exact budget after sorting",
			costs: []int{4, 1, 2},
			coins: 3,
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIceCream(tt.costs, tt.coins); got != tt.want {
				t.Errorf("maxIceCream() = %v, want %v", got, tt.want)
			}
		})
	}
}
