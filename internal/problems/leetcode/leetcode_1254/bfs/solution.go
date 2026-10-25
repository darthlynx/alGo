package bfs

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
			if grid[i][j] == 0 && !visited[i][j] && bfs(i, j, grid, visited) {
				islands++
			}
		}
	}

	return islands
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func bfs(i int, j int, grid [][]int, visited [][]bool) bool {
	queue := [][]int{}
	head := 0

	queue = append(queue, []int{i, j})
	visited[i][j] = true

	isClosed := true

	for head < len(queue) { // while queue is not empty
		current := queue[head]
		head++

		for k := range 4 {
			x := current[0] + dx[k]
			y := current[1] + dy[k]

			if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
				isClosed = false
			} else if grid[x][y] == 0 && !visited[x][y] {
				queue = append(queue, []int{x, y})
				visited[x][y] = true
			}
		}
	}

	return isClosed
}
