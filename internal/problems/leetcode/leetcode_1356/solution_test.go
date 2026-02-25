package leetcode1356

import (
	"reflect"
	"testing"
)

func TestSortByBits(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			"Example 1",
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			[]int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			"Example 2",
			[]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			[]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByBits(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByBits(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}
