package leetcode167

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// TwoSum finds two numbers in a sorted array that add up to the target
// Input array is 1-indexed and the solution must use constant extra space
//
// Time Complexity: O(n log n)
// Space Complexity: O(1)
func TwoSum(numbers []int, target int) []int {
	for i, v := range numbers {
		index := binarySearch(numbers[i+1:], target-v)
		if index >= 0 {
			return []int{i + 1, i + index + 2} // +1 for 1-based indexing, +1 for slicing offset
		}
	}
	return []int{-1, -1}
}

// binarySearch performs binary search in a sorted array
func binarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
