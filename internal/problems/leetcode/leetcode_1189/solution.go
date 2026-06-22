package leetcode1189

import "math"

// https://leetcode.com/problems/maximum-number-of-balloons/
//
// Time Complexity: O(n).
// Space Complexity: O(1).
func maxNumberOfBalloons(text string) int {
	freq := make(map[rune]int, 26)

	for _, ch := range text {
		freq[ch] = freq[ch] + 1
	}

	balloon := "balloon"

	minFreq := math.MaxInt
	for _, ch := range balloon {
		val, ok := freq[ch]
		if !ok {
			return 0
		}
		if (ch == 'l' || ch == 'o') && val < 2 {
			return 0
		}
		minFreq = min(minFreq, val)
	}
	frL := freq['l']
	frO := freq['o']
	minFreq = min(minFreq, frL/2)
	minFreq = min(minFreq, frO/2)

	return minFreq
}
