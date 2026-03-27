package leetcode2946

// https://leetcode.com/problems/matrix-similarity-after-cyclic-shifts/
//
// Time Complexity: O(m*n).
// Space Complexity: O(1).
func areSimilar(mat [][]int, k int) bool {
	m := len(mat)
	n := len(mat[0])
	k %= n // reduce the k to perform less operations

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if r%2 == 0 {
				// check left shift
				// (c-k+n)%n works only if k is reduced by n, otherwise it will cause out of bound error
				if mat[r][(c-k+n)%n] != mat[r][c] {
					return false
				}
			} else {
				// check right shift
				if mat[r][(c+k)%n] != mat[r][c] {
					return false
				}
			}
		}
	}

	return true
}
