package leetcode678

// https://leetcode.com/problems/valid-parenthesis-string/
//
// Time complexity: O(n)
// Space complexity: O(1)
func checkValidString(s string) bool {
	open := 0
	close := 0

	// checking pairs from left and right
	// two pointers approach
	end := len(s) - 1
	for start := 0; start < len(s); start++ {
		if s[start] == '(' || s[start] == '*' {
			open++
		} else {
			open--
		}

		if s[end] == ')' || s[end] == '*' {
			close++
		} else {
			close--
		}

		// means that string is unbalanced
		if open < 0 || close < 0 {
			return false
		}
		end--
	}
	return true
}
