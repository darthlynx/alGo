package leetcode121

import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock
func maxProfit(prices []int) int {
	minSum := math.MaxInt32
	maxProfit := 0

	for i := range prices {
		if prices[i] < minSum {
			minSum = prices[i]
		} else {
			if prices[i]-minSum > maxProfit {
				maxProfit = prices[i] - minSum
			}
		}
	}

	return maxProfit
}
