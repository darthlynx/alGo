package leetcode2833

// https://leetcode.com/problems/furthest-point-from-origin/
//
// Time complexity: O(n).
// Space complexity: O(1).
func furthestDistanceFromOrigin(moves string) int {
	pos := 0
	unCount := 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'L':
			pos--
		case 'R':
			pos++
		default:
			unCount++
		}
	}
	return abs(pos) + unCount
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
