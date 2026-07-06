package leetcode1288

import "slices"

// https://leetcode.com/problems/remove-covered-intervals/
//
// Time complexity: O(n*log(n))
// Space complexity: O(log(n)) for sorting
func removeCoveredIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] == b[0] {
			return b[1] - a[1]
		}
		return a[0] - b[0]
	})

	toRemove := 0
	bestRightBoundary := intervals[0][1]
	toCompare := 0
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		prev := intervals[toCompare]
		if curr[0] >= prev[0] && curr[1] <= prev[1] {
			toRemove++
		} else {
			if curr[1] > bestRightBoundary {
				bestRightBoundary = curr[1]
				toCompare = i
			}
		}
	}

	return len(intervals) - toRemove
}
