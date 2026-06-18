package leetcode1344

import (
	"math"
	"testing"
)

func TestAngleClock(t *testing.T) {
	tests := []struct {
		name    string
		hour    int
		minutes int
		want    float64
	}{
		{"12:30", 12, 30, 165},
		{"3:30", 3, 30, 75},
		{"3:15", 3, 15, 7.5},
		{"4:50", 4, 50, 155},
		{"12:00", 12, 0, 0},
		{"6:00", 6, 0, 180},
		{"1:00", 1, 0, 30},
		{"1:57", 1, 57, 76.5},
		{"12:59", 12, 59, 35.5},
		{"9:00", 9, 0, 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := angleClock(tt.hour, tt.minutes); math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("angleClock() = %v, want %v", got, tt.want)
			}
		})
	}
}
