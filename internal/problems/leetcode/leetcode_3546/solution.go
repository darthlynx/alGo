package leetcode3546

// https://leetcode.com/problems/equal-sum-grid-partition-i/
//
// Time Complexity: O(m*n).
// Space Complexity: O(1).
func canPartitionGrid(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])

	// total sum
	total := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			total += grid[r][c]
		}
	}

	// check all horizontal cuts
	topSum := 0
	for r := 0; r < m-1; r++ {
		for c := 0; c < n; c++ {
			topSum += grid[r][c]
		}
		if topSum*2 == total {
			return true
		}
	}

	// check all vertical cuts
	leftSum := 0
	for c := 0; c < n-1; c++ {
		for r := 0; r < m; r++ {
			leftSum += grid[r][c]
		}
		if leftSum*2 == total {
			return true
		}
	}

	return false
}
