package leetcode1339


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

const modulo = 1000000007

// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
//
// Time Complexity: O(n)
// Space Complexity: O(h), where h is the height of the tree
func maxProduct(root *TreeNode) int {
    totalSum := totalSumDfs(root)
    maxProductVal := 0
    dfs(root, totalSum, &maxProductVal)
    return maxProductVal % modulo
}

func dfs(root *TreeNode, totalSum int, maxProductVal *int) int {
    if root == nil {
        return 0
    }
    left := dfs(root.Left, totalSum, maxProductVal)
    right := dfs(root.Right, totalSum, maxProductVal)
    subtreeSum := root.Val + left + right

    product := (totalSum - subtreeSum) * subtreeSum
    if  product > *maxProductVal {
        *maxProductVal = product
    }
    return subtreeSum
}


func totalSumDfs(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return root.Val + totalSumDfs(root.Left) + totalSumDfs(root.Right)
}
