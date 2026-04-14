package leetcode1848

import "testing"

func TestGetMinDistance(t *testing.T) {
	type args struct {
		nums   []int
		target int
		start  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 5,
				start:  3,
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				nums:   []int{1},
				target: 1,
				start:  0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinDistance(tt.args.nums, tt.args.target, tt.args.start); got != tt.want {
				t.Errorf("getMinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
