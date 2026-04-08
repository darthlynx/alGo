package leetcode3653

import "testing"

func TestXorAfterQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				nums:    []int{1, 1, 1},
				queries: [][]int{{0, 2, 1, 4}},
			},
			want: 4,
		},
		{
			name: "test case 2",
			args: args{
				nums:    []int{2, 3, 1, 5, 4},
				queries: [][]int{{1, 4, 2, 3}, {0, 2, 1, 2}},
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorAfterQueries(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("xorAfterQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
