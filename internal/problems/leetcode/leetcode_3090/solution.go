package leetcode3090

// https://leetcode.com/problems/maximum-length-substring-with-two-occurrences/
//
// Time Complexity: O(n), where n is the length of the input string
// Space Complexity: O(1)
func maximumLengthSubstring(s string) int {
	var charCount [26]int
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		charCount[s[right]-'a']++

		for left <= right && charCount[s[right]-'a'] > 2 {
			charCount[s[left]-'a']--
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
