package leetcode1758

// https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/
//
// Time Complexity: O(n) - where n is the length of the input string
// Space Complexity: O(1)
func minOperations(s string) int {
	mismatch := 0
	// Generating pattern which starts with 0
	// 01010...
	// 0 % 2 = 0
	// 1 % 2 = 1
	// 2 % 2 = 0
	// 3 % 2 = 1 ...
	for i := 0; i < len(s); i++ {
		expected := byte('0')
		if i%2 == 1 {
			expected = byte('1')
		}

		if s[i] != expected {
			mismatch++
		}
	}

	// checking other pattern 101010
	// if position matches pattern 01010 then it mismatches 10101
	otherMismatch := len(s) - mismatch

	return min(mismatch, otherMismatch)
}
