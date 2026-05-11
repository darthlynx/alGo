package leetcode2553

import (
	"reflect"
	"testing"
)

func TestSeparateDigits(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example 1: mixed multi-digit numbers", []int{13, 25, 83, 77}, []int{1, 3, 2, 5, 8, 3, 7, 7}},
		{"example 2: numbers with repeated digits", []int{7, 1, 3, 9}, []int{7, 1, 3, 9}},
		{"single element single digit", []int{5}, []int{5}},
		{"single element multi-digit", []int{123}, []int{1, 2, 3}},
		{"all single-digit numbers", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"large numbers", []int{100, 999}, []int{1, 0, 0, 9, 9, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := separateDigits(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("separateDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
