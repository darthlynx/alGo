package leetcode238

// https://leetcode.com/problems/product-of-array-except-self/
//
// Time Complexity: O(n).
// Space Complexity: O(n) or O(1) if not including the output array.
func productExceptSelf(nums []int) []int {
    n := len(nums)
    product := make([]int, n)

    product[0] = 1
    for i := 1; i < n; i++ {
        product[i] = product[i-1] * nums[i-1]
    }

    suffix := 1
    for i := n-1; i >= 0; i-- {
        product[i] = product[i] * suffix
        suffix = suffix * nums[i]
    }

    return product
}
