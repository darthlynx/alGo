package leetcode1967

import "strings"

// https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/
//
// Time Complexity: O(n*m), where n is the length of patterns and m is the length of word.
// Space Complexity: O(1).
func numOfStrings(patterns []string, word string) int {
	count := 0

	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			count++
		}
	}

	return count
}
