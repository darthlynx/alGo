package leetcode657

// https://leetcode.com/problems/robot-return-to-origin/
//
// Time Complexity: O(n), where n is the length of `moves`.
// Space Complexity: O(1).
func judgeCircle(moves string) bool {
	leftRight := 0
	upDown := 0

	movesTypes := map[byte]int{
		'R': 1,
		'L': -1,
		'U': 1,
		'D': -1,
	}

	for i := 0; i < len(moves); i++ {
		move := moves[i]
		if move == 'R' || move == 'L' {
			leftRight += movesTypes[move]
		} else {
			upDown += movesTypes[move]
		}
	}

	return leftRight == 0 && upDown == 0
}
