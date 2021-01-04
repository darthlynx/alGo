package main

import (
	"fmt"
	"sort"
)

func main() {
	salary := []int{8000,9000,2000,3000,6000,1000}
	fmt.Println(salary)
	fmt.Println(average(salary))
	fmt.Println(salary)
}

// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
func average(salary []int) float64 {
	sorted := salary
	sort.Ints(sorted)
	sorted = sorted[1 : len(sorted)-1]
	var sum float64
	for _, v := range sorted {
		sum += float64(v)
	}
	return sum / float64(len(salary)-2)
}