package leetcode657

// https://leetcode.com/problems/robot-return-to-origin/
//
// Time Complexity: O(n), where n is the length of `moves`.
// Space Complexity: O(1).
func judgeCircle(moves string) bool {
	leftRight := 0
	upDown := 0

	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'L':
			leftRight++
		case 'R':
			leftRight--
		case 'U':
			upDown++
		case 'D':
			upDown--
		}
	}

	return leftRight == 0 && upDown == 0
}
