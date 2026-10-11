package leetcode1541

// https://leetcode.com/problems/minimum-insertions-to-balance-a-parentheses-string/
//
// Time complexity: O(n)
// Space complexity: O(1)
func minInsertions(s string) int {
	toAdd := 0
	open := 0

	for i := 0; i < len(s); i++{
		ch := s[i]
		if ch == '(' {
			open++
		} else {
			// stack is empty, need to add an opening bracket
			if open == 0 {
				toAdd++
			}
			// have two closing brackets, moving extra step ahead
			if i+1 < len(s) && s[i+1] == ')' {
				i++
			} else { // need one more )
				toAdd++
			}
			if open > 0 {
				open--
			}
		}
	}

	return toAdd + open*2
}
