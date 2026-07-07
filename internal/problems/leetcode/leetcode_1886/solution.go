package leetcode1886

// https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/
//
// Time Complexity: O(n^2).
// Space Complexity: O(1).
func findRotation(mat [][]int, target [][]int) bool {
	for i := 0; i < 4; i++ {
		if equal(mat, target) {
			return true
		}
		mat = rotate(mat)
	}
	return false
}

func rotate(target [][]int) [][]int {
	target = transpose(target) // makes rows as columns
	return reflect(target)     // reverse each row to get the rotated matrix
}

func transpose(target [][]int) [][]int {
	n := len(target)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp := target[i][j]
			target[i][j] = target[j][i]
			target[j][i] = tmp
		}
	}
	return target
}

func reflect(target [][]int) [][]int {
	n := len(target)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			tmp := target[i][j]
			target[i][j] = target[i][n-j-1]
			target[i][n-j-1] = tmp
		}
	}
	return target
}

func equal(mat [][]int, target [][]int) bool {
	n := len(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}
	return true
}
