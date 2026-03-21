package leetcode3643

import (
	"reflect"
	"testing"
)

func TestReverseSubmatrix(t *testing.T) {
	type args struct {
		grid [][]int
		x    int
		y    int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test-case-1",
			args: args{
				grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				x:    0,
				y:    0,
				k:    2,
			},
			want: [][]int{{4, 5, 3}, {1, 2, 6}, {7, 8, 9}},
		},
		{
			name: "test-case-2",
			args: args{
				grid: [][]int{{1, 2}, {3, 4}},
				x:    0,
				y:    0,
				k:    2,
			},
			want: [][]int{{3, 4}, {1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseSubmatrix(tt.args.grid, tt.args.x, tt.args.y, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
