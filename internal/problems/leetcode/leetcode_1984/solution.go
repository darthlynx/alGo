package leetcode1984

import "sort"

//https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
//
// Time Complexity: O(n log n) where n is the length of nums
// Space Complexity: O(1)
func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)
    minDiff := nums[len(nums)-1] - nums[0]

    for i := 0; i <= len(nums)-k; i++ {
        diff := nums[i + k - 1] - nums[i]
        if diff < minDiff {
            minDiff = diff
        }
    }
    return minDiff
}
