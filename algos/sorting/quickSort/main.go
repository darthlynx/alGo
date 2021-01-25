package main

import (
	"fmt"
)

func main() {
	ii := []int{10, 5, 2, 3}
	fmt.Println(ii)
	fmt.Println(quickSort(ii))
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var less []int
	var greater []int

	pivotIndex := 0
	pivot := arr[pivotIndex]

	for _, v := range arr[pivotIndex+1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	var result []int
	result = append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}
