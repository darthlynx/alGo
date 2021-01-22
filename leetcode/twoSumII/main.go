package main

import (
	"fmt"
)

func main() {
	xx := []int{2, 7, 11, 15}
	target := 9 // Expected result [1,2]
	fmt.Println(twoSum(xx, target))
	xx = []int{0, 0, 3, 4}
	target = 0 // Expected result [1,2]
	fmt.Println(twoSum(xx, target))
}

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) []int {
	for i, v := range numbers {
		index := UpperBound(numbers, target-v)
		fmt.Printf("elem index: %v; index: %v\n", i, index)
		if index >= 0 {
			return []int{i + 1, index + 1}
		}
	}
	return []int{-1, -1}
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
