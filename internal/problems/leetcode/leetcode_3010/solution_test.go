package leetcode3010

import "testing"


func TestMinimumCost(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "one",
			nums: []int{1,2,3,12},
			want: 6,
		},
		{
			name: "two",
			nums: []int{5,4,3},
			want: 12,
		},
		{
			name: "three",
			nums: []int{10,3,1,1},
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumCost(tt.nums)
			if got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
