package leetcode796

import "strings"

// https://leetcode.com/problems/rotate-string/
//
// Time complexity: O(n) due to string concatenation and search.
// Space complexity: O(n) for the doubled string.
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	doubled := s + s
	return strings.Contains(doubled, goal)
}
