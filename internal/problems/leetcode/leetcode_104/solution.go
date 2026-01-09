package leetcode104


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
//
// Time Complexity: O(n)
// Space Complexity: O(h), where h is the height of the tree
func maxDepth(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(root *TreeNode, depth int) int {
    if root == nil {
        return depth
    }

    return 1 + max(dfs(root.Left, depth), dfs(root.Right, depth))
}
