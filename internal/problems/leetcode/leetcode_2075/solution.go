package leetcode2075

import "strings"

// https://leetcode.com/problems/decode-the-slanted-ciphertext/
//
// Time Complexity: O(n), where n is the length of `encodedText`.
// Space Complexity: O(n).
func decodeCiphertext(encodedText string, rows int) string {
	original := make([]byte, 0, len(encodedText))
	n := len(encodedText) / rows
	for i := 0; i < n; i++ {
		idx := i
		for idx < len(encodedText) {
			original = append(original, encodedText[idx])
			idx += (n + 1)
		}
	}
	return strings.TrimRight(string(original), " ")
}
