package leetcode1559

import "testing"

func TestContainsCycle(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want bool
	}{
		{
			name: "example 1 - a forms outer cycle",
			grid: [][]byte{
				{'a', 'a', 'a', 'a'},
				{'a', 'b', 'b', 'a'},
				{'a', 'b', 'b', 'a'},
				{'a', 'a', 'a', 'a'},
			},
			want: true,
		},
		{
			name: "example 2 - c forms cycle",
			grid: [][]byte{
				{'c', 'c', 'c', 'a'},
				{'c', 'd', 'c', 'c'},
				{'c', 'c', 'e', 'c'},
				{'f', 'c', 'c', 'c'},
			},
			want: true,
		},
		{
			name: "example 3 - no cycle",
			grid: [][]byte{
				{'a', 'b', 'b'},
				{'b', 'z', 'b'},
				{'b', 'b', 'a'},
			},
			want: false,
		},
		{
			name: "2x2 all same - minimal cycle",
			grid: [][]byte{
				{'a', 'a'},
				{'a', 'a'},
			},
			want: true,
		},
		{
			name: "single row - no cycle possible",
			grid: [][]byte{
				{'a', 'a', 'a'},
			},
			want: false,
		},
		{
			name: "single cell - no cycle",
			grid: [][]byte{
				{'a'},
			},
			want: false,
		},
		{
			name: "all different characters - no cycle",
			grid: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsCycle(tt.grid)
			if got != tt.want {
				t.Errorf("containsCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
