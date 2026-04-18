package leetcode3783

import "testing"

func TestMirrorDistance(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"Example 1", 123, 198},
		{"Example 2", 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorDistance(tt.n); got != tt.want {
				t.Errorf("mirrorDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
