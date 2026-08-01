package dp

// https://leetcode.com/problems/predict-the-winner/
//
// Time Complexity: O(n^2)
// Space Complexity: O(n^2)
func predictTheWinner(nums []int) bool {
	n := len(nums)
	memo := make([][]int, n)
	visited := make([][]bool, n)
	for i := range n {
		memo[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}

	return dp(0, len(nums)-1, nums, memo, visited) >= 0
}

func dp(left int, right int, nums []int, memo [][]int, visited [][]bool) int {
	if left == right {
		return nums[left]
	}
	if visited[left][right] {
		return memo[left][right]
	}
	route1 := nums[left] - dp(left+1, right, nums, memo, visited)
	route2 := nums[right] - dp(left, right-1, nums, memo, visited)
	res := max(route1, route2)
	memo[left][right] = res
	visited[left][right] = true
	return res
}
