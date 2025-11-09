package leetcode1491

import "testing"

func TestAverage(t *testing.T) {
	tests := []struct {
		name   string
		salary []int
		want   float64
	}{
		{
			name:   "example from problem",
			salary: []int{4000, 3000, 1000, 2000},
			want:   2500.00000,
		},
		{
			name:   "another example",
			salary: []int{1000, 2000, 3000},
			want:   2000.00000,
		},
		{
			name:   "all same salaries",
			salary: []int{5000, 5000, 5000, 5000},
			want:   5000.00000,
		},
		{
			name:   "large range of salaries",
			salary: []int{6000, 5000, 4000, 3000, 2000, 1000},
			want:   3500.00000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := average(tt.salary)
			if got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
