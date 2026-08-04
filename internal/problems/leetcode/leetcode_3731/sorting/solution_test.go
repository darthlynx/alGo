package sorting

import (
	"slices"
	"testing"
)

func TestFindMissingElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example 1", []int{1, 4, 2, 5}, []int{3}},
		{"example 2", []int{7, 8, 6, 9}, []int{}},
		{"example 3", []int{5, 1}, []int{2, 3, 4}},
		{"single element", []int{5}, []int{}},
		{"consecutive already sorted", []int{3, 4, 5}, []int{}},
		{"multiple gaps", []int{1, 3, 6, 7, 10}, []int{2, 4, 5, 8, 9}},
		{"negative range", []int{-3, 1}, []int{-2, -1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingElements(tt.nums); !slices.Equal(got, tt.want) {
				t.Errorf("findMissingElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
