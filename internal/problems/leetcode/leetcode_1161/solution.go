package leetcode1161

import "math"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func maxLevelSum(root *TreeNode) int {
    var queue []*TreeNode
    head := 0
    queue = append(queue, root)

    level := 1
    maxSumLevel := level
    maxSum := math.MinInt

    // while "queue" is not empty
    for (head < len(queue)) {
        size := len(queue) - head
        var levelSum int
        for i := 0; i < size; i++ {
            node := queue[head]
            head++
            levelSum += node.Val

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        if levelSum > maxSum {
            maxSumLevel = level
            maxSum = levelSum
        }
        level++
    }

    return maxSumLevel
}
