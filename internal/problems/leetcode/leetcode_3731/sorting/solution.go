package sorting

import "slices"

// https://leetcode.com/problems/find-missing-elements/
//
// Time Complexity: O(n logn + m), where m is the range between smallest and biggest numbers)
// Space Complexity: O(n + m)
func findMissingElements(nums []int) []int {
	var res []int
	slices.Sort(nums)
	count := nums[0]
	i := 0
	for i < len(nums) {
		if nums[i] == count {
			i++
		} else if nums[i] > count {
			res = append(res, count)
		}
		count++
	}

	return res
}
