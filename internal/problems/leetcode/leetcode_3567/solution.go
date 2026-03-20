package leetcode3567

import (
	"math"
	"slices"
)

// https://leetcode.com/problems/minimum-absolute-difference-in-sliding-submatrix/
//
// Time Complexity: O((m-k+1)*(n-k+1)* k^2 log k).
// Space Complexity: O(k^2).
func minAbsDiff(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	rows := m - k + 1
	cols := n - k + 1

	result := make([][]int, rows)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			minAbsDiff := math.MaxInt
			dist := make(map[int]struct{})
			for i := row; i < row+k; i++ {
				for j := col; j < col+k; j++ {
					dist[grid[i][j]] = struct{}{}
				}
			}
			nums := []int{}
			for key := range dist {
				nums = append(nums, key)
			}
			slices.Sort(nums)
			for i := 0; i < len(nums)-1; i++ {
				if abs(nums[i]-nums[i+1]) < minAbsDiff {
					minAbsDiff = abs(nums[i] - nums[i+1])
				}
			}
			if minAbsDiff == math.MaxInt {
				minAbsDiff = 0
			}
			result[row] = append(result[row], minAbsDiff)
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
