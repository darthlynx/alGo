package leetcode2161

// https://leetcode.com/problems/partition-array-according-to-given-pivot/
//
// Time Complexity: O(n),
// Space Complexity: O(n).
func pivotArray(nums []int, pivot int) []int {
	smaller := make([]int, 0, len(nums))
	bigger := make([]int, 0, len(nums))

	n := len(nums)
	for _, num := range nums {
		if num < pivot {
			smaller = append(smaller, num)
		} else if num > pivot {
			bigger = append(bigger, num)
		}
	}

	k := n - len(smaller) - len(bigger)

	for range k {
		smaller = append(smaller, pivot)
	}

	smaller = append(smaller, bigger...)
	return smaller
}
