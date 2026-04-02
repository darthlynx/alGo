package leetcode3418

import "math"

// https://leetcode.com/problems/maximum-amount-of-money-robot-can-earn/
//
// Time Complexity: O(m*n).
// Space Complexity: O(m*n).
func maximumAmount(coins [][]int) int {
	m := len(coins)
	n := len(coins[0])
	const nr = 2 // how many robbers can be neutralized

	dp := make([][][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, nr+1)
			for k := 0; k <= nr; k++ {
				dp[i][j][k] = math.MinInt
			}
		}
	}

	// initialize first cell
	// case 1: do not neutralize
	dp[0][0][nr] = coins[0][0]
	// case 2: neutralize
	if coins[0][0] < 0 && nr > 0 {
		dp[0][0][nr-1] = 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for k := 0; k <= nr; k++ {
				// from top
				if i > 0 {
					// case 1: do not neutralize here
					if dp[i-1][j][k] != math.MinInt {
						dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k]+coins[i][j])
					}

					// case 2: neutralize here
					if k+1 <= nr && dp[i-1][j][k+1] != math.MinInt {
						candidate := dp[i-1][j][k+1] + max(coins[i][j], 0)
						dp[i][j][k] = max(dp[i][j][k], candidate)
					}
				}

				// from left
				if j > 0 {
					// case 1: do not neutralize here
					if dp[i][j-1][k] != math.MinInt {
						dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k]+coins[i][j])
					}

					// case 2: neutralize here
					if k+1 <= nr && dp[i][j-1][k+1] != math.MinInt {
						candidate := dp[i][j-1][k+1] + max(coins[i][j], 0)
						dp[i][j][k] = max(dp[i][j][k], candidate)
					}
				}
			}
		}
	}

	result := math.MinInt
	for k := 0; k <= nr; k++ {
		result = max(result, dp[m-1][n-1][k])
	}
	return result
}
