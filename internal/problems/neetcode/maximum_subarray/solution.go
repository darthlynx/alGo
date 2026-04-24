package maximumsubarray

import "math"

// https://neetcode.io/problems/maximum-subarray/
//
// Time complexity: O(n).
// Space complexity: O(1).
func maxSubArray(nums []int) int {
	// Kadane's algorithm
	maxSum := math.MinInt
	currentSum := 0
	for _, num := range nums {
		currentSum = max(num, num+currentSum)
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}
