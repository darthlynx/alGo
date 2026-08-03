package leetcode1406

import "math"

// https://leetcode.com/problems/stone-game-iii/
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		best := math.MinInt
		sum := 0
		for k := 0; k < 3 && k+i < n; k++ {
			sum += stoneValue[i+k]
			best = max(best, sum-dp[k+i+1])
		}

		dp[i] = best
	}

	if dp[0] > 0 {
		return "Alice"
	}
	if dp[0] < 0 {
		return "Bob"
	}
	return "Tie"
}
