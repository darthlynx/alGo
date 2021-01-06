package main

import "fmt"

func main() {
	fmt.Println("hello")
}

// https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	return nil
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
