package leetcode3567

import (
	"reflect"
	"testing"
)

func TestMinAbsDiff(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want [][]int
	}{
		{"Example 1", [][]int{{1, 8}, {3, -2}}, 2, [][]int{{2}}},
		{"Example 2", [][]int{{3, -1}}, 1, [][]int{{0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAbsDiff(tt.grid, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minAbsDiff(%v, %v) = %v, want %v", tt.grid, tt.k, got, tt.want)
			}
		})
	}
}
