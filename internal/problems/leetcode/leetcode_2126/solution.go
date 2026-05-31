package leetcode2126

import "slices"

// https://leetcode.com/problems/destroying-asteroids/
//
// Time complexity: O(n log n),
// Space complexity: O(log n) due to sorting
func asteroidsDestroyed(mass int, asteroids []int) bool {
	slices.Sort(asteroids)

	currentMass := mass

	for _, asteroid := range asteroids {
		if asteroid > currentMass {
			return false
		}
		currentMass += asteroid
	}
	return true
}
