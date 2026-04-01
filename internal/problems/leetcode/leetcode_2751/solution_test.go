package leetcode2751

import (
	"reflect"
	"testing"
)

func TestSurvivedRobotsHealths(t *testing.T) {
	tests := []struct {
		positions  []int
		healths    []int
		directions string
		want       []int
	}{
		{
			positions:  []int{5, 4, 3, 2, 1},
			healths:    []int{2, 17, 9, 15, 10},
			directions: "RRRRR",
			want:       []int{2, 17, 9, 15, 10},
		},
		{
			positions:  []int{3, 5, 2, 6},
			healths:    []int{10, 10, 15, 12},
			directions: "RLRL",
			want:       []int{14},
		},
		{
			positions:  []int{1, 2, 5, 6},
			healths:    []int{10, 10, 11, 11},
			directions: "RLRL",
			want:       []int{},
		},
		{
			positions:  []int{3, 40},
			healths:    []int{49, 11},
			directions: "LL",
			want:       []int{49, 11},
		},
		{
			positions:  []int{11, 44, 16},
			healths:    []int{1, 20, 17},
			directions: "RLR",
			want:       []int{18},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := survivedRobotsHealths(tt.positions, tt.healths, tt.directions)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
