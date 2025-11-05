package main

import "fmt"

func main() {
	ii := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(ii, target))
	jj := []int{5, 7, 7, 8, 8, 10}
	target = 6
	fmt.Println(searchRange(jj, target))
	var kk []int
	target = 0
	fmt.Println(searchRange(kk, target))
	ll := []int{1}
	fmt.Println(searchRange(ll, target))
}

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	right := UpperBound(nums, target)
	left := LowerBound(nums, target)
	return []int{left, right}
}

func LowerBound(arr []int, target int) int {
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

func UpperBound(arr []int, target int) int {
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
