package main

import "fmt"

func main() {

	prices := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(prices))
}

const IntMax = int(^uint(0) >> 1)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock
func maxProfit(prices []int) int {
	minSum := IntMax
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minSum {
			minSum = prices[i]
		} else {
			maxProfit = maxOf(maxProfit, prices[i] - minSum)
		}
	}

	return maxProfit
}

func maxOf(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
