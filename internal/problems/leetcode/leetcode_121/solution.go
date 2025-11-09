package leetcode121

import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock
func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}

		currentProfit := prices[i] - minPrice
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
	}

	return maxProfit
}
