package leetcode2033

import "slices"

// https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/
//
// Time complexity: O(m*n*log(m*n)) due to sorting.
// Space complexity: O(m*n) for the line slice.
func minOperations(grid [][]int, x int) int {
	m := len(grid)
	n := len(grid[0])

	line := make([]int, m*n)
	k := 0
	for i := range m {
		for j := range n {
			line[k] = grid[i][j]
			k++
		}
	}

	slices.Sort(line)

	median := line[len(line)/2]

	count := 0
	for _, num := range line {
		if num%x != median%x {
			return -1
		}
		count += abs(num-median) / x
	}
	return count
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
