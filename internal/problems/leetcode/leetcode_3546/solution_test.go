package leetcode3546

import "testing"

func TestCanPartitionGrid(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-case-1",
			args: args{
				grid: [][]int{{1, 4}, {2, 3}},
			},
			want: true,
		},
		{
			name: "test-case-2",
			args: args{
				grid: [][]int{{1, 3}, {2, 4}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionGrid(tt.args.grid); got != tt.want {
				t.Errorf("canPartitionGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
