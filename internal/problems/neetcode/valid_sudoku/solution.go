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
	for row := 0; row < n; row += 3 {
		for col := 0; col < n; col += 3 {

			seen := make(map[byte]bool, 9)
			for i := row; i < row+3; i++ {
				for j := col; j < col+3; j++ {
					v := board[i][j]
					if v == '.' {
						continue
					}
					if seen[v] {
						return false
					}
					seen[v] = true
				}
			}
		}
	}

	return true
}

func isValidRow(i int, board [][]byte) bool {
	seen := make(map[byte]bool, 9)
	for j := 0; j < len(board); j++ {
		v := board[i][j]
		if v == '.' {
			continue
		}
		if seen[v] { // found duplicate
			return false
		}
		seen[v] = true
	}
	return true
}

func isValidCol(j int, board [][]byte) bool {
	seen := make(map[byte]bool, 9)
	for i := 0; i < len(board); i++ {
		v := board[i][j]
		if v == '.' {
			continue
		}
		if seen[v] { // found duplicate
			return false
		}
		seen[v] = true
	}
	return true
}
