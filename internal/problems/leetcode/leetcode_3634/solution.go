package leetcode3634

import "sort"

// https://leetcode.com/problems/minimum-removals-to-balance-array/
//
// Time Complexity: O(n log n)
// Space Complexity: O(log n) - due to sorting
func minRemoval(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    minRm := n
    right := 0
    for left := 0; left < n; left++ {
        for right < n && nums[right] <= k * nums[left] {
            right++
        }
        minRm = min(minRm, n - (right - left))
    }
    return minRm
}
