package leetcode1861

import (
	"reflect"
	"testing"
)

func TestRotateTheBox(t *testing.T) {
	tests := []struct {
		name    string
		boxGrid [][]byte
		want    [][]byte
	}{
		{
			name:    "single row stones with gap",
			boxGrid: [][]byte{{'#', '.', '#'}},
			want:    [][]byte{{'.'}, {'#'}, {'#'}},
		},
		{
			name:    "two rows with obstacles and stones",
			boxGrid: [][]byte{{'#', '.', '*', '.'}, {'#', '#', '*', '#'}},
			want:    [][]byte{{'#', '.'}, {'#', '#'}, {'*', '*'}, {'#', '.'}},
		},
		{
			name:    "all empty",
			boxGrid: [][]byte{{'.', '.', '.'}},
			want:    [][]byte{{'.'}, {'.'}, {'.'}},
		},
		{
			name:    "all stones",
			boxGrid: [][]byte{{'#', '#', '#'}},
			want:    [][]byte{{'#'}, {'#'}, {'#'}},
		},
		{
			name:    "stone blocked above obstacle",
			boxGrid: [][]byte{{'#', '*', '.'}},
			want:    [][]byte{{'#'}, {'*'}, {'.'}},
		},
		{
			name:    "stone falls below obstacle",
			boxGrid: [][]byte{{'*', '#', '.'}},
			want:    [][]byte{{'*'}, {'.'}, {'#'}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateTheBox(tt.boxGrid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateTheBox() = %v, want %v", got, tt.want)
			}
		})
	}
}
