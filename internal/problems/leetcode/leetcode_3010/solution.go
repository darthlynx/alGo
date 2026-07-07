package leetcode3010

import (
	"sort"
)

// https://leetcode.com/problems/divide-an-array-into-subarrays-with-minimum-cost-i/
//
// Time Complexity: O(n log n)
// Space Complexity: O(1)
func minimumCost(nums []int) int {
	first := nums[0]
	rem := nums[1:]
	sort.Ints(rem)
	return first + rem[0] + rem[1]
}
