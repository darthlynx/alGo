package leetcode1727

import "slices"

// https://leetcode.com/problems/largest-submatrix-with-rearrangements/
//
// Time complexity: O(m * n * log(n)).
// Space complexity: O(1).
func largestSubmatrix(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row > 0 && matrix[row][col] != 0 {
				matrix[row][col] += matrix[row-1][col]
			}
		}
	}

	result := 0
	for _, row := range matrix {
		slices.SortFunc(row, func(a, b int) int {
			return b - a
		})
		for i := 0; i < n; i++ {
			result = max(result, row[i]*(i+1))
		}
	}

	return result
}
