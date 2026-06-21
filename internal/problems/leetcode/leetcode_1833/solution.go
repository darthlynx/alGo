package leetcode1833

import "slices"

// https://leetcode.com/problems/maximum-ice-cream-bars/
//
// Time Complexity: O(n*log n).
// Space Complexity: O(1).
func maxIceCream(costs []int, coins int) int {
	slices.Sort(costs)

	count := 0
	for _, cost := range costs {
		if cost > coins {
			break
		}
		count++
		coins -= cost
	}
	return count
}
