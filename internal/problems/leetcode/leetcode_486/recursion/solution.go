package recursion

// https://leetcode.com/problems/predict-the-winner/
//
// Time Complexity: O(2^n)
// Space Complexity: O(n)
func predictTheWinner(nums []int) bool {
	return dp(0, len(nums)-1, nums) >= 0
}

func dp(left int, right int, nums []int) int {
	if left == right {
		return nums[left]
	}
	route1 := nums[left] - dp(left+1, right, nums)
	route2 := nums[right] - dp(left, right-1, nums)
	return max(route1, route2)
}
