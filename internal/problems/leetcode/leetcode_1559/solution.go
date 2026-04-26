package leetcode1559

var (
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

// https://leetcode.com/problems/detect-cycles-in-2d-grid/
//
// Time complexity: O(m*n).
// Space complexity: O(m*n).
func containsCycle(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])

	// track visited points
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}

	for i := range m {
		for j := range n {
			// skip visited
			if visited[i][j] {
				continue
			}
			if hasCycle(grid[i][j], [2]int{i, j}, grid, visited, i, j) {
				return true
			}
		}
	}

	return false
}

func hasCycle(root byte, prev [2]int, grid [][]byte, visited [][]bool, i, j int) bool {
	// out of bounds
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return false
	}

	ch := grid[i][j]
	current := [2]int{i, j}

	if ch != root {
		return false
	}

	if visited[i][j] {
		return true // found the cycle
	}

	// mark as visited
	visited[i][j] = true

	res := false
	for k := range 4 {
		nX := i + dx[k]
		nY := j + dy[k]
		// skip parent
		if nX == prev[0] && nY == prev[1] {
			continue
		}
		res = res || hasCycle(root, current, grid, visited, nX, nY)
	}
	return res
}
