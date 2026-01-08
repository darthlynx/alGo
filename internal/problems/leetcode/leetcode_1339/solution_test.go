package leetcode1339

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
			want: 110,
		},
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: 90,
		},
	}

	for _, tc := range tests {
		got := maxProduct(tc.root)
		if got != tc.want {
			t.Errorf("maxProduct(%v) = %v; want %v", tc.root, got, tc.want)
		}
	}
}
