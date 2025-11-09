package leetcode53

import "math"

// https://leetcode.com/problems/maximum-subarray/
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func maxSubArray(nums []int) int {

	maxSubarray := math.MinInt32
	currentSum := maxSubarray

	for _, val := range nums {
		// updating current subarray
		if currentSum+val > val {
			currentSum += val
		} else {
			currentSum = val
		}

		if currentSum > maxSubarray {
			maxSubarray = currentSum
		}
	}

	return maxSubarray
}
