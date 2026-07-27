package leetcode628

import "slices"

// https://leetcode.com/problems/maximum-product-of-three-numbers/
//
// Time Complexity: O(n * log(n))
// Space Complexity: O(log(n)) due to sorting
func maximumProduct(nums []int) int {
	n := len(nums)

	slices.Sort(nums)

	// negative path
	way1 := nums[0] * nums[1] * nums[n-1]
	// positive path
	way2 := nums[n-1] * nums[n-2] * nums[n-3]

	return max(way1, way2)
}
