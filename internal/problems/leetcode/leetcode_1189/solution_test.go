package leetcode1189

import "testing"

func TestMaxNumberOfBalloons(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{"example 1", "nlaebolko", 1},
		{"example 2", "loonbalxballpoon", 2},
		{"example 3", "leetcode", 0},
		{"empty string", "", 0},
		{"exactly one balloon", "balloon", 1},
		{"missing letter b", "alloonalloon", 0},
		{"only one l", "balon", 0},
		{"no o", "balln", 0},
		{"three balloons", "balloonballoonballoon", 3},
		{"excess l and o, single b limits", "balllllloooooon", 1},
		{"single l, double o", "baloon", 0},
		{"extra l and o, b is bottleneck", "balloonloon", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfBalloons(tt.text); got != tt.want {
				t.Errorf("maxNumberOfBalloons() = %v, want %v", got, tt.want)
			}
		})
	}
}
