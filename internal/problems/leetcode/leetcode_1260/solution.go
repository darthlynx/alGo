package leetcode1260

// https://leetcode.com/problems/shift-2d-grid/
//
// Time Complexity: O(m * n)
// Space Complexity: O(m * n)
func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	total := m * n
	k %= total

	flat := make([]int, total)
	i := 0

	for row := range m {
		for col := range n {
			flat[i] = grid[row][col]
			i++
		}
	}

	result := make([][]int, m)
	for row := range m {
		result[row] = make([]int, n)
	}

	for i := range total {
		newPos := (i + k) % total
		col := newPos % n
		row := newPos / n
		result[row][col] = flat[i]
	}
	return result
}
