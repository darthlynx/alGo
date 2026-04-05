package leetcode657

import "testing"

func TestJudgeCircle(t *testing.T) {
	type args struct {
		moves string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-case-1",
			args: args{
				moves: "UD",
			},
			want: true,
		},
		{
			name: "test-case-2",
			args: args{
				moves: "LL",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeCircle(tt.args.moves); got != tt.want {
				t.Errorf("judgeCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
