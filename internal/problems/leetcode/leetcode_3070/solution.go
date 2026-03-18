package leetcode3070

// https://leetcode.com/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/
//
// Time complexity: O(m * n).
// Space complexity: O(m * n).
func countSubmatrices(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	prefixSum := make([][]int, m)
	result := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res := grid[i][j]
			if j > 0 {
				res += prefixSum[i][j-1]
			}
			if i > 0 {
				res += prefixSum[i-1][j]
			}
			if i > 0 && j > 0 {
				// remove overlap, because top-left area
				// belongs to both prefixSum[i][j-1] and prefixSum[i-1][j]
				res -= prefixSum[i-1][j-1]
			}
			if res <= k {
				result++
			}
			prefixSum[i] = append(prefixSum[i], res)
		}
	}
	return result
}
