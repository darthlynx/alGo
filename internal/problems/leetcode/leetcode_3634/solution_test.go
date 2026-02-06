package leetcode3634

import "testing"

func TestMinRemoval(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "first example",
			nums: []int{2, 1, 5},
			k:    2,
			want: 1,
		},
		{
			name: "second example",
			nums: []int{1, 6, 2, 9},
			k:    3,
			want: 2,
		},
		{
			name: "third example",
			nums: []int{4, 6},
			k:    2,
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := minRemoval(test.nums, test.k)
			if got != test.want {
				t.Errorf("minRemoval(%v, %d) = %d; want %d", test.nums, test.k, got, test.want)
			}
		})
	}
}
