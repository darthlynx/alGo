package leetcode34

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	right := upperBound(nums, target)
	left := lowerBound(nums, target)
	return []int{left, right}
}

func lowerBound(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right) / 2
		if target <= arr[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if arr[left] == target {
		return left
	}
	return -1
}

func upperBound(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right + 1) / 2
		if target >= arr[mid] {
			left = mid
		} else {
			right = mid - 1
		}
	}
	if arr[right] == target {
		return right
	}
	return -1
}
