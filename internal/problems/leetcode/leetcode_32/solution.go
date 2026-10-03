package leetcode_32

// https://leetcode.com/problems/longest-valid-parentheses/
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func longestValidParentheses(s string) int {
	longest := 0
	score := 0
	// if len(s) < 2 {
	// 	return 0
	// }

	left := 0
	for right := 0; right < len(s); right++ {
		ch := s[right]
		if ch == '(' {
			score++
		} else {
			score--
		}
		if score == 0 {
			longest = max(longest, right-left+1)
		}
		if score < 0 {
			left = right + 1
			score = 0
		}
	}

	// reversed order
	score = 0 // reset
	right := len(s) - 1
	for left := len(s) - 1; left >= 0; left-- {
		ch := s[left]
		if ch == ')' {
			score++
		} else {
			score--
		}
		if score == 0 {
			longest = max(longest, right-left+1)
		}
		if score < 0 {
			right = left - 1
			score = 0
		}
	}

	return longest
}
