package leetcode2075

import "testing"

func TestDecodeCiphertext(t *testing.T) {
	type args struct {
		encodedText string
		rows        int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-case-1",
			args: args{
				encodedText: "ch   ie   pr",
				rows:        3,
			},
			want: "cipher",
		},
		{
			name: "test-case-2",
			args: args{
				encodedText: "iveo    eed   l te   olc",
				rows:        4,
			},
			want: "i love leetcode",
		},
		{
			name: "test-case-3",
			args: args{
				encodedText: " b  ac",
				rows:        2,
			},
			want: " abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeCiphertext(tt.args.encodedText, tt.args.rows); got != tt.want {
				t.Errorf("decodeCiphertext() = %v, want %v", got, tt.want)
			}
		})
	}
}
