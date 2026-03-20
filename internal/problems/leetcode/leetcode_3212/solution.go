package leetcode3212

// https://leetcode.com/problems/count-submatrices-with-equal-frequency-of-x-and-y/
//
// Time Complexity: O(m * n).
// Space Complexity: O(m * n).
func numberOfSubmatrices(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	transform := map[byte]int{
		'X': 1,
		'Y': -1,
		'.': 0,
	}

	count := 0
	prefixSum := make([][]int, m)
	xCount := make([][]int, m)

	for row := 0; row < m; row++ {
		prefixSum[row] = make([]int, n)
		xCount[row] = make([]int, n)
		for col := 0; col < n; col++ {
			prefixSum[row][col] = transform[grid[row][col]]
			if grid[row][col] == 'X' {
				xCount[row][col] = 1
			}
			if col > 0 {
				prefixSum[row][col] += prefixSum[row][col-1]
				xCount[row][col] += xCount[row][col-1]
			}
			if row > 0 {
				prefixSum[row][col] += prefixSum[row-1][col]
				xCount[row][col] += xCount[row-1][col]
			}
			if row > 0 && col > 0 {
				prefixSum[row][col] -= prefixSum[row-1][col-1]
				xCount[row][col] -= xCount[row-1][col-1]
			}
			if prefixSum[row][col] == 0 && xCount[row][col] > 0 {
				count++
			}
		}
	}

	return count
}
