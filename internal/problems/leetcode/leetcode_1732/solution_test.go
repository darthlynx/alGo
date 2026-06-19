package leetcode1732

import "testing"

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		name string
		gain []int
		want int
	}{
		{
			name: "example 1",
			gain: []int{-5, 1, 5, 0, -7},
			want: 1,
		},
		{
			name: "example 2",
			gain: []int{-4, -3, -2, -1, 4, 3, 2},
			want: 0,
		},
		{
			name: "single positive gain",
			gain: []int{5},
			want: 5,
		},
		{
			name: "single negative gain",
			gain: []int{-5},
			want: 0,
		},
		{
			name: "all positive gains",
			gain: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "all negative gains",
			gain: []int{-1, -2, -3},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestAltitude(tt.gain); got != tt.want {
				t.Errorf("largestAltitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
