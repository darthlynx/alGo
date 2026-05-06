package leetcode1861

const (
	stone    = '#'
	obstacle = '*'
	empty    = '.'
)

// https://leetcode.com/problems/rotating-the-box/
//
// Time complexity: O(m*n).
// Space complexity: O(m*n).
func rotateTheBox(boxGrid [][]byte) [][]byte {

	m := len(boxGrid)
	n := len(boxGrid[0])

	for i := range m {
		emptyPos := n - 1
		for j := n - 1; j >= 0; j-- {
			switch boxGrid[i][j] {
			case obstacle:
				emptyPos = j - 1
			case stone:
				boxGrid[i][j] = empty
				boxGrid[i][emptyPos] = stone
				emptyPos--
			}
		}
	}

	rotatedBox := make([][]byte, n)
	for i := range n {
		rotatedBox[i] = make([]byte, m)
		for j := range m {
			rotatedBox[i][j] = boxGrid[m-1-j][i]
		}
	}

	return rotatedBox
}
