package leetcode3643

// https://leetcode.com/problems/flip-square-submatrix-vertically/
//
// Time Complexity: O(k^2) ~> O(m*n).
// Space Complexity: O(1).
func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	for col := y; col < y+k; col++ {
		top := x
		bottom := x + k - 1
		for top < bottom {
			tmp := grid[top][col]
			grid[top][col] = grid[bottom][col]
			grid[bottom][col] = tmp
			top++
			bottom--
		}
	}
	return grid
}
