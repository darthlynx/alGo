package leetcode868

import (
	"strconv"
)

// https://leetcode.com/problems/binary-gap/
//
// Time Complexity: O(log n) - where n is the input number
// Space Complexity: O(log n) - due to binary string representation
func binaryGap(n int) int {
	s := strconv.FormatInt(int64(n), 2) // to binary format
	longest := 0

	left := 0
	for left < len(s) {
		if s[left] == '1' {
			right := left + 1
			for right < len(s) && s[right] != '1' {
				right++
			}
			if right < len(s) && s[right] == '1' {
				longest = max(longest, right-left)
			}
			left = right
		} else {
			left++
		}
	}

	return longest
}
