package leetcode2196

import (
	"reflect"
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	tests := []struct {
		name         string
		descriptions [][]int
		want         *TreeNode
	}{
		{
			name:         "example 1",
			descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}},
			want: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 17},
				},
				Right: &TreeNode{
					Val:  80,
					Left: &TreeNode{Val: 19},
				},
			},
		},
		{
			name:         "example 2",
			descriptions: [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
		},
		{
			name:         "single left edge",
			descriptions: [][]int{{1, 2, 1}},
			want: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
		},
		{
			name:         "single right edge",
			descriptions: [][]int{{5, 3, 0}},
			want: &TreeNode{
				Val:   5,
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name:         "left chain",
			descriptions: [][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
		},
		{
			name:         "root not first in input",
			descriptions: [][]int{{2, 4, 0}, {1, 2, 1}, {1, 3, 0}},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createBinaryTree(tt.descriptions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
