package leetcode1861

// https://leetcode.com/problems/rotating-the-box/
//
// Time complexity: O(m*n).
// Space complexity: O(m*n).
func rotateTheBox(boxGrid [][]byte) [][]byte {
	const stone = '#'
	const obstacle = '*'
	const empty = '.'

	m := len(boxGrid)
	n := len(boxGrid[0])

	rotatedBox := make([][]byte, n)
	for i := range n {
		rotatedBox[i] = make([]byte, m)
	}

	for i := range n {
		for j := range m {
			rotatedBox[i][j] = boxGrid[m-1-j][i]
		}
	}

	for col := range m {
		bottom := n - 1
		top := n - 1
		for bottom >= 0 {
			for bottom >= 0 && rotatedBox[bottom][col] != empty {
				bottom--
				top = bottom - 1
			}

			if top < 0 {
				break
			}

			if rotatedBox[top][col] == obstacle {
				top--
				bottom = top
				continue
			}
			if rotatedBox[top][col] == empty {
				top--
				continue
			}
			if rotatedBox[top][col] == stone && rotatedBox[bottom][col] == empty {
				rotatedBox[bottom][col] = stone
				rotatedBox[top][col] = empty
				bottom--
				top--
				continue
			}
		}
	}

	return rotatedBox
}
