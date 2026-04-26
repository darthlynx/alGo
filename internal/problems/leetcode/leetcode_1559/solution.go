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
	visited := make(map[[2]int]struct{})

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// skip visited
			if _, ok := visited[[2]int{i, j}]; ok {
				continue
			}
			if hasCycle(grid[i][j], [2]int{i, j}, grid, visited, i, j) {
				return true
			}
		}
	}

	return false
}

func hasCycle(root byte, prev [2]int, grid [][]byte, visited map[[2]int]struct{}, i, j int) bool {
	// out of bounds
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return false
	}

	ch := grid[i][j]
	current := [2]int{i, j}

	if ch != root {
		return false
	}

	if _, ok := visited[current]; ok {
		return true // found the cycle
	}

	// mark as visited
	visited[current] = struct{}{}

	res := false
	for k := 0; k < 4; k++ {
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
