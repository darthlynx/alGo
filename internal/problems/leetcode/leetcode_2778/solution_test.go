package leetcode2778

import "testing"

func TestSumOfSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example with four elements",
			nums: []int{1, 2, 3, 4},
			want: 21,
		},
		{
			name: "example with six elements",
			nums: []int{2, 7, 1, 19, 18, 3},
			want: 63,
		},
		{
			name: "empty input",
			nums: []int{},
			want: 0,
		},
		{
			name: "single minimum value",
			nums: []int{1},
			want: 1,
		},
		{
			name: "prime length excludes interior elements",
			nums: []int{2, 50, 50, 50, 3},
			want: 13,
		},
		{
			name: "all same values with multiple divisors",
			nums: []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
			want: 54,
		},
		{
			name: "maximum length and values",
			nums: []int{
				50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
				50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
				50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
				50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
				50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
			},
			want: 15000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfSquares(tt.nums); got != tt.want {
				t.Errorf("sumOfSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
