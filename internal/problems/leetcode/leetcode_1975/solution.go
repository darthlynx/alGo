package leetcode1975

import "math"

// https://leetcode.com/problems/maximum-matrix-sum/
//
// Time Complexity: O(m * n)
// Space Complexity: O(1)
func maxMatrixSum(matrix [][]int) int64 {
	m := len(matrix)
	n := len(matrix[0])
	var maxSum int64
	negatives := 0
	absSum := 0
	smallestNegative := math.MaxInt
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] < 0 {
				negatives++
			}
			absVal := abs(matrix[i][j])
			if absVal < smallestNegative {
				smallestNegative = absVal
			}
			absSum += absVal
		}
	}
	if negatives%2 == 0 {
		maxSum += int64(absSum)
	} else {
		// Subtract twice the smallest absolute value
		// (once to remove it from sum, once to make it negative)
		maxSum += int64(absSum - 2*smallestNegative)
	}
	return maxSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
