package leetcode865

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/
//
// Time Complexity: O(n)
// Space Complexity: O(h), where h is the height of the tree
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	node, _ := dfs(root, 0)
	return node
}

// returns common ancestor for the deepest node and max depth for the subtree
func dfs(root *TreeNode, depth int) (*TreeNode, int) {
	if root == nil {
		return nil, depth
	}

	left, leftDepth := dfs(root.Left, depth+1)
	right, rightDepth := dfs(root.Right, depth+1)

	// if both left and right have same depth, then current node is the common ancestor
	if leftDepth == rightDepth {
		return root, leftDepth
	}

	if leftDepth > rightDepth {
		return left, leftDepth
	}
	return right, rightDepth
}
