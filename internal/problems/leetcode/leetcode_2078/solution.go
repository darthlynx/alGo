package leetcode2078

// https://leetcode.com/problems/two-furthest-houses-with-different-colors/
//
// Time complexity: O(n).
// Space complexity: O(1).
func maxDistance(colors []int) int {
	n := len(colors)
	maxDist := 0

	// pin rightmost
	for i := 0; i < n; i++ {
		if colors[i] == colors[n-1] {
			continue
		} else {
			maxDist = max(maxDist, n-1-i)
			break
		}
	}

	// pin leftmost
	for i := n - 1; i >= 0; i-- {
		if colors[i] == colors[0] {
			continue
		} else {
			maxDist = max(maxDist, i)
			break
		}
	}

	return maxDist
}
