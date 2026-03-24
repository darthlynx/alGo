package leetcode2906

// https://leetcode.com/problems/construct-product-matrix/
//
// Time Complexity: O(m*n).
// Space Complexity: O(m*n) or O(1) if not including the output array.
func constructProductMatrix(grid [][]int) [][]int {
	const modulo = 12345
	m := len(grid)
	n := len(grid[0])

	p := make([][]int, m)
	for row := range m {
		p[row] = make([]int, n)
	}

	// store prefix production in the result array
	prefix := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p[i][j] = prefix
			prefix = (prefix * (grid[i][j] % modulo)) % modulo
		}
	}

	// update result array with the suffix production
	suffix := 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			p[i][j] = (p[i][j] * suffix) % modulo
			suffix = (suffix * (grid[i][j] % modulo)) % modulo
		}
	}

	return p
}
