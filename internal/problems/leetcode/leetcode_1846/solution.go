package leetcode1846

import "slices"

// https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/
//
// Time Complexity: O(n*logn).
// Space Complexity: O(1).
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	slices.Sort(arr)
	maxVal := 1

	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
		maxVal = max(maxVal, arr[i])
	}
	return maxVal
}
