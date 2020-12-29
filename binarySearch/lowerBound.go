package main

import "fmt"

func main() {
	// items := []int{2, 3, 7, 8, 10}
	items := []int{}
	fmt.Println("Lower bound", LowerBound(items, 4))
	fmt.Println("Upper bound", UpperBound(items, 4))
}

func LowerBound(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	min := 0
	max := len(arr) - 1
	for min < max {
		mid := (max + min) / 2
		if target <= arr[mid] {
			max = mid
		} else {
			min = mid + 1
		}
	}
	if arr[min] < target {
		return -1
	} else {
		return arr[min]
	}
}

func UpperBound(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	min := 0
	max := len(arr) - 1
	for min < max {
		mid := (min + max + 1) / 2
		if target >= arr[mid] {
			min = mid
		} else {
			max = mid - 1
		}
	}
	if arr[max] > target {
		return -1
	} else {
		return arr[max]
	}
}
