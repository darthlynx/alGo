package dfs

// https://leetcode.com/problems/number-of-closed-islands/
//
// Time Complexity: O(m * n)
// Space Complexity: O(m * n)
func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}

	islands := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && !visited[i][j] && dfs(i, j, grid, visited) {
				islands++
			}
		}
	}

	return islands
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func dfs(i int, j int, grid [][]int, visited [][]bool) bool {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return false // not a closed island
	}

	if grid[i][j] == 1 || visited[i][j] {
		return true
	}

	visited[i][j] = true
	isClosed := true

	for k := range 4 {
		x := i + dx[k]
		y := j + dy[k]
		if !dfs(x, y, grid, visited) {
			isClosed = false
		}
	}

	return isClosed
}
