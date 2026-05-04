package leetcode48

// https://leetcode.com/problems/rotate-image/
//
// Time complexity: O(n^2).
// Space complexity: O(1).
func rotate(matrix [][]int) {
	transpose(matrix)
	mirror(matrix)
}

func transpose(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func mirror(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			k := n - j - 1
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}
