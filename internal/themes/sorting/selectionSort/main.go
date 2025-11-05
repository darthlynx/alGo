package main

import (
	"fmt"
)

func main() {
	xx := []int{5, 3, 6, 2, 10, 1}
	fmt.Println("Initial:", xx)
	xx = selectionSort(xx)
	fmt.Println("Result:", xx)
}

func findSmallestElemIndex(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i, v := range arr {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}
	return smallestIndex
}

// O(n^2)
func selectionSort(arr []int) []int {
	var resArr []int
	for len(arr) > 0 {
		smallestElemIndex := findSmallestElemIndex(arr) // find the index of the smallest element
		resArr = append(resArr, arr[smallestElemIndex]) // put it into result slice
		deleteElem(&arr, smallestElemIndex)             // remove smallest element from initial slice
	}
	return resArr
}

func deleteElem(arr *[]int, i int) {
	left := (*arr)[:i]
	right := (*arr)[i+1:]
	*arr = append(left, right...)
}
