package leetcode3016

import "slices"

// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-ii/
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func minimumPushes(word string) int {
	numOfKeys := 8

	freq := make([]int, 26)
	for _, ch := range word {
		freq[ch-'a']++
	}

	slices.SortFunc(freq, func(a, b int) int {
		return b - a
	})

	result := 0
	for i := range 26 {
		if freq[i] == 0 {
			break
		}
		result += (i/numOfKeys + 1) * freq[i]
	}

	return result
}
