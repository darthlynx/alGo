package leetcode2265

import "testing"

func TestAverageOfSubtree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "example from problem",
			root: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}},
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 6},
				},
			},
			want: 5,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: 1,
		},
		{
			name: "nil root",
			root: nil,
			want: 0,
		},
		{
			name: "all same values",
			root: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}},
			want: 3,
		},
		{
			name: "integer division truncation makes root match",
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4}},
			want: 3,
		},
		{
			name: "root differs from its subtree average",
			root: &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfSubtree(tt.root); got != tt.want {
				t.Errorf("averageOfSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
