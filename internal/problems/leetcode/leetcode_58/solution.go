package leetcode58

// https://leetcode.com/problems/length-of-last-word/
//
// Time Complexity: O(n) - where n is the length of the input string
// Space Complexity: O(1)
func lengthOfLastWord(s string) int {
	i := len(s) - 1

	// skip trailing spaces
	for i >= 0 && s[i] == ' ' {
		i--
	}

	size := 0
	for i >= 0 && s[i] != ' ' {
		size++
		i--
	}
	return size
}
