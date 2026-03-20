package leetcode3212

import "testing"

func TestNumberOfSubmatrices(t *testing.T) {
	tests := []struct {
		name string
		nums [][]byte
		want int
	}{
		{
			name: "first example",
			nums: [][]byte{{'X', 'Y', '.'}, {'Y', '.', '.'}},
			want: 3,
		},
		{
			name: "second example",
			nums: [][]byte{{'X', 'X'}, {'X', 'Y'}},
			want: 0,
		},
		{
			name: "third example",
			nums: [][]byte{{'.', '.'}, {'.', '.'}},
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := numberOfSubmatrices(test.nums)
			if got != test.want {
				t.Errorf("numberOfSubmatrices(%v) = %d; want %d", test.nums, got, test.want)
			}
		})
	}
}
