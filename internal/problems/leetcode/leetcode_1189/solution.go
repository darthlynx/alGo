package leetcode1189

import "math"

// https://leetcode.com/problems/maximum-number-of-balloons/
//
// Time Complexity: O(n).
// Space Complexity: O(1).
func maxNumberOfBalloons(text string) int {
	freq := [26]int{}
	for _, ch := range text {
		freq[ch-'a']++
	}

	// 'balloon' has double l and o, so we need to divide their frequency by 2
	minFreq := math.MaxInt
	for _, ch := range "ban" {
		minFreq = min(minFreq, freq[ch-'a'])
	}
	for _, ch := range "lo" {
		minFreq = min(minFreq, freq[ch-'a']/2)
	}

	return minFreq
}
