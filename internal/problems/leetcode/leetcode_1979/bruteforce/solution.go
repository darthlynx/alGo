package bruteforce

import "math"

// https://leetcode.com/problems/find-greatest-common-divisor-of-array/
//
// Time Complexity: O(n+mn)
// Space Complexity: O(1)
func findGCD(nums []int) int {
	mn := math.MaxInt
	mx := math.MinInt

	for _, num := range nums {
		mn = min(mn, num)
		mx = max(mx, num)
	}

	greatestCommon := 1
	for i := 1; i <= mn; i++ {
		if mn%i == 0 && mx%i == 0 {
			greatestCommon = max(greatestCommon, i)
		}
	}

	return greatestCommon
}
