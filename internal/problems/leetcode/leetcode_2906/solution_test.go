package leetcode2906

import (
	"reflect"
	"testing"
)

func TestConstructProductMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test-case-1",
			args: args{
				grid: [][]int{{1, 2}, {3, 4}},
			},
			want: [][]int{{24, 12}, {8, 6}},
		},
		{
			name: "test-case-2",
			args: args{
				grid: [][]int{{12345}, {2}, {1}},
			},
			want: [][]int{{2}, {0}, {0}},
		},
		{
			name: "test-case-3",
			args: args{
				grid: [][]int{{12345, 12345}},
			},
			want: [][]int{{0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructProductMatrix(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructProductMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
