package leetcode3546

// https://leetcode.com/problems/equal-sum-grid-partition-i/
//
// Time Complexity: O(m*n).
// Space Complexity: O(m+n).
func canPartitionGrid(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])

	total := 0
	// sum per row
	rows := make([]int, m)
	for i := 0; i < m; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum += grid[i][j]
		}
		rows[i] = sum
		total += sum
	}

	// check all horizontal cuts
	currentSum := 0
	for r := 0; r < m; r++ {
		currentSum += rows[r]
		if total-currentSum == currentSum {
			return true
		}
	}

	// calculate columns sums
	cols := make([]int, n)
	for j := 0; j < n; j++ {
		sum := 0
		for i := 0; i < m; i++ {
			sum += grid[i][j]
		}
		cols[j] = sum
	}

	// check all vertical cuts
	currentSum = 0
	for c := 0; c < n; c++ {
		currentSum += cols[c]
		if total-currentSum == currentSum {
			return true
		}
	}

	return false
}
