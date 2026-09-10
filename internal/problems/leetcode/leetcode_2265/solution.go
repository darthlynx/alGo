package leetcode2265

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/
//
// Time Complexity: O(n), where n is the number of nodes in the tree
// Space Complexity: O(h), where h is the height of the tree
func averageOfSubtree(root *TreeNode) int {
	_, _, result := dfs(root)
	return result
}

func dfs(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0
	}

	sumLeft, cntLeft, resL := dfs(root.Left)
	sumRight, cntRight, resR := dfs(root.Right)

	sum := sumLeft + sumRight + root.Val
	count := cntLeft + cntRight + 1

	average := sum / count

	res := resL + resR

	if average == root.Val {
		res++
	}

	return sum, count, res
}
