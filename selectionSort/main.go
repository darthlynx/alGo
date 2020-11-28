package main

import (
	"fmt"
)

func main() {
	xx := []int{5, 3, 6, 2, 10}
	fmt.Println("Initial:", xx)
	fmt.Println("Result:", selectionSort(xx))
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i, _ := range arr {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

// O(n^2)
func selectionSort(arr []int) []int {
	var resArr []int
	for len(arr) > 0 {
		smallestIndex := findSmallest(arr)                          // find the index of the smallest element
		resArr = append(resArr, arr[smallestIndex])                 // put it into result slice
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...) // remove smallest element from initial slice
	}
	return resArr
}
