package main

import "fmt"

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func main() {
	nums := []int {-2,1,-3,4,-1,2,1,-5,4}
	expectedResult := 6

	result := maxSubArray(nums)
	fmt.Printf("expected: %v; actual: %v\n", expectedResult, result)
}

func maxSubArray(nums []int) int {
	maxSum := MinInt
	currentSum := 0

	for _,v := range nums {
		currentSum = maxOf(v, currentSum + v)
		maxSum = maxOf(maxSum, currentSum)
	}

	return maxSum
}


func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}