package leetcode1848

// https://leetcode.com/problems/minimum-distance-to-the-target-element/
//
// Time complexity: O(n).
// Space complexity: O(1).
func getMinDistance(nums []int, target int, start int) int {
	minDist := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			minDist = min(minDist, abs(i-start))
		}
	}
	return minDist
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
