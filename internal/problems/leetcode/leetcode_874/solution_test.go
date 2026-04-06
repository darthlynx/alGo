package leetcode874

import "testing"

func TestRobotSim(t *testing.T) {
	tests := []struct {
		name      string
		commands  []int
		obstacles [][]int
		want      int
	}{
		{
			name:      "test case 1",
			commands:  []int{4, -1, 3},
			obstacles: [][]int{},
			want:      25,
		},
		{
			name:      "test case 2",
			commands:  []int{4, -1, 4, -2, 4},
			obstacles: [][]int{{2, 4}},
			want:      65,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotSim(tt.commands, tt.obstacles); got != tt.want {
				t.Errorf("robotSim() = %v, want %v", got, tt.want)
			}
		})
	}
}
