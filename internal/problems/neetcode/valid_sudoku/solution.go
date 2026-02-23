package validsudoku

// https://neetcode.io/problems/valid-sudoku/question
//
// Time Complexity: O(n^2) - where n is the size of the board (9 in this case)
// Space Complexity: O(n) - due to the hash maps used for validation
func isValidSudoku(board [][]byte) bool {
	n := len(board)

	for i := 0; i < n; i++ {
		if !isValidRow(i, board) {
			return false
		}
	}

	for j := 0; j < n; j++ {
		if !isValidCol(j, board) {
			return false
		}
	}

	// check squares
	row := 0
	col := 0
	for row < n {

		for col < n {
			m := make(map[byte]bool, 9)
			for i := row; i < row+3; i++ {
				for j := col; j < col+3; j++ {
					if _, ok := m[board[i][j]]; ok {
						return false
					}
					if board[i][j] != '.' {
						m[board[i][j]] = true
					}
				}
			}
			col += 3
		}
		row += 3
		col = 0
	}

	return true
}

func isValidRow(i int, board [][]byte) bool {
	m := make(map[byte]bool, 9)
	for j := 0; j < len(board); j++ {
		if _, ok := m[board[i][j]]; ok { // found duplicate
			return false
		}
		if board[i][j] != '.' {
			m[board[i][j]] = true
		}
	}
	return true
}

func isValidCol(j int, board [][]byte) bool {
	m := make(map[byte]bool, 9)
	for i := 0; i < len(board); i++ {
		if _, ok := m[board[i][j]]; ok { // found duplicate
			return false
		}
		if board[i][j] != '.' {
			m[board[i][j]] = true
		}
	}
	return true
}
