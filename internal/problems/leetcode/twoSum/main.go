package main

import "fmt"

func main() {
	xx := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(xx, target))
}

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	checked := make(map[int]int)

	for i, el := range nums {
		if v, ok := checked[target-el]; ok {
			return []int{i, v}
		} else {
			checked[el] = i
		}
	}
	return []int{-1, -1}
}
