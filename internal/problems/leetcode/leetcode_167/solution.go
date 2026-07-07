package leetcode167

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for right-left > 0 {
		currentSum := numbers[left] + numbers[right]
		if currentSum == target {
			break
		}
		if currentSum > target {
			right--
		} else {
			left++
		}
	}

	return []int{left + 1, right + 1} // 1-based indexing
}
