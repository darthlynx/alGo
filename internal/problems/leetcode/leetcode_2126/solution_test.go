package leetcode2126

import "testing"

func TestAsteroidsDestroyed(t *testing.T) {
	tests := []struct {
		name      string
		mass      int
		asteroids []int
		want      bool
	}{
		{"example 1 from problem", 10, []int{3, 9, 19, 5, 21}, true},
		{"example 2 from problem", 5, []int{4, 9, 23, 4}, false},
		{"empty asteroids", 7, []int{}, true},
		{"single asteroid equal to mass", 5, []int{5}, true},
		{"single asteroid greater than mass", 5, []int{6}, false},
		{"sort is load-bearing", 5, []int{4, 9, 8}, true},
		{"largest asteroid first would fail unsorted", 2, []int{3, 1}, true},
		{"large values within constraints", 100000, []int{100000, 100000, 100000}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidsDestroyed(tt.mass, tt.asteroids); got != tt.want {
				t.Errorf("asteroidsDestroyed() = %v, want %v", got, tt.want)
			}
		})
	}
}
