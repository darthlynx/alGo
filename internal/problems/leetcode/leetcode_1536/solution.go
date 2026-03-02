package leetcode1536

// https://leetcode.com/problems/minimum-swaps-to-arrange-a-binary-grid/
//
// Time Complexity: O(n^2) - where n is the size of the grid
// Space Complexity: O(n) - due to the rightmostOne array
func minSwaps(grid [][]int) int {
	n := len(grid)

	rightmostOne := make([]int, n)
	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				rightmostOne[i] = j
				break
			}
		}
	}

	swaps := 0
	for i := 0; i < n; i++ {
		k := -1
		for j := i; j < n; j++ {
			if rightmostOne[j] <= i {
				swaps += j - i
				k = j
				break
			}
		}
		if k < 0 {
			return -1
		}
		for j := k; j > i; j-- {
			temp := rightmostOne[j]
			rightmostOne[j] = rightmostOne[j-1]
			rightmostOne[j-1] = temp
		}
	}
	return swaps
}
