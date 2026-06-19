package leetcode1732

// https://leetcode.com/problems/find-the-highest-altitude/
//
// Time Complexity: O(n).
// Space Complexity: O(1).
func largestAltitude(gain []int) int {
	highest := 0
	altitude := 0

	for _, g := range gain {
		altitude += g
		highest = max(highest, altitude)
	}

	return highest
}
