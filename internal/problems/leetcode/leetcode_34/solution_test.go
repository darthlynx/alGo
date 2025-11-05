package leetcode34

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "found target with multiple occurrences",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "target not found",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "empty array",
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
		{
			name:   "single element array, target not found",
			nums:   []int{1},
			target: 0,
			want:   []int{-1, -1},
		},
		{
			name:   "single element array, target found",
			nums:   []int{1},
			target: 1,
			want:   []int{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRange(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerBound(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{
			name:   "found target at start",
			arr:    []int{1, 1, 2, 3},
			target: 1,
			want:   0,
		},
		{
			name:   "found target in middle",
			arr:    []int{1, 2, 2, 3},
			target: 2,
			want:   1,
		},
		{
			name:   "target not found",
			arr:    []int{1, 2, 3, 4},
			target: 5,
			want:   -1,
		},
		{
			name:   "empty array",
			arr:    []int{},
			target: 1,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowerBound(tt.arr, tt.target)
			if got != tt.want {
				t.Errorf("lowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{
			name:   "found target at end",
			arr:    []int{1, 2, 3, 3},
			target: 3,
			want:   3,
		},
		{
			name:   "found target in middle",
			arr:    []int{1, 2, 2, 3},
			target: 2,
			want:   2,
		},
		{
			name:   "target not found",
			arr:    []int{1, 2, 3, 4},
			target: 5,
			want:   -1,
		},
		{
			name:   "empty array",
			arr:    []int{},
			target: 1,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := upperBound(tt.arr, tt.target)
			if got != tt.want {
				t.Errorf("upperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
