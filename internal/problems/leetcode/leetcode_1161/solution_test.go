package leetcode1161

import "testing"

func TestMaxLevelSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: -8,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			want: 2,
		},
		{
			root: &TreeNode{
				Val: 989,
				Right: &TreeNode{
					Val: 10250,
					Left: &TreeNode{
						Val: 98693,
					},
					Right: &TreeNode{
						Val: -89388,
						Right: &TreeNode{
							Val: -32127,
						},
					},
				},
			},
			want: 2,
		},
	}

	for _, tc := range tests {
		got := maxLevelSum(tc.root)
		if got != tc.want {
			t.Errorf("maxLevelSum(%v) = %v; want %v", tc.root, got, tc.want)
		}
	}
}
