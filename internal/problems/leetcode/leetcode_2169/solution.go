package leetcode2169

// https://leetcode.com/problems/count-operations-to-obtain-zero/
//
// Time Complexity: O(max(num1, num2))
// Space Complexity: O(1)
func countOperations(num1 int, num2 int) int {
	count := 0

	for num1 > 0 && num2 > 0 {
		if num1 >= num2 {
			num1 = num1 - num2
		} else {
			num2 = num2 - num1
		}
		count++
	}

	return count
}
