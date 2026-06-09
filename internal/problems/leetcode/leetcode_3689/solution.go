package leetcode3689

import "math"

// https://leetcode.com/problems/maximum-total-subarray-value-i/
//
// Time Complexity: O(n),
// Space Complexity: O(1).
func maxTotalValue(nums []int, k int) int64 {
	mmin := math.MaxInt
	mmax := math.MinInt

	for _, num := range nums {
		mmin = min(mmin, num)
		mmax = max(mmax, num)
	}
	// we could choose the whole array k times
	return int64(mmax-mmin) * int64(k)
}
