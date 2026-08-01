package recursion

import "testing"

func TestPredictTheWinner(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"example 1", []int{1, 5, 2}, false},
		{"example 2", []int{1, 5, 233, 7}, true},
		{"single element", []int{5}, true},
		{"two elements", []int{1, 2}, true},
		{"all same even count", []int{2, 2, 2, 2}, true},
		{"all same odd count", []int{3, 3, 3}, true},
		{"large ends dominate", []int{20, 30, 2, 2, 2, 10}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictTheWinner(tt.nums); got != tt.want {
				t.Errorf("predictTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
