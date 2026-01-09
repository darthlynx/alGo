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
	maxDepth := getMaxDepth(root, 0)
	return dfs(root, 0, maxDepth)
}

func dfs(root *TreeNode, depth int, maxDepth int) *TreeNode {
	if root == nil {
		return nil
	}

	if depth+1 == maxDepth {
		return root
	}

	left := dfs(root.Left, depth+1, maxDepth)
	right := dfs(root.Right, depth+1, maxDepth)

	// if both left and right are not nil, it means this is the common ancestor
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

// DFS to get max depth of the tree
func getMaxDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	return 1 + max(getMaxDepth(root.Left, depth), getMaxDepth(root.Right, depth))
}
