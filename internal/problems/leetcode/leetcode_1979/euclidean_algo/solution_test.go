package euclideanalgo

import "testing"

func TestFindGCD(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{2, 5, 6, 9, 10}, 2},
		{"example 2", []int{7, 5, 6, 8, 3}, 1},
		{"example 3", []int{3, 3}, 3},
		{"single element", []int{12}, 12},
		{"min divides max", []int{4, 8, 16}, 4},
		{"min ones", []int{1, 1000}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGCD(tt.nums); got != tt.want {
				t.Errorf("findGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
