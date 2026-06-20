package leetcode1840

import (
	"cmp"
	"slices"
)

// https://leetcode.com/problems/maximum-building-height/
//
// Time Complexity: O(n*logn).
// Space Complexity: O(n).
func maxBuilding(n int, restrictions [][]int) int {
	restrictions = append(restrictions, []int{1, 0})
	slices.SortFunc(restrictions, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	if restrictions[len(restrictions)-1][0] != n {
		restrictions = append(restrictions, []int{n, n - 1})
	}

	m := len(restrictions)
	// from left to right
	for i := 1; i < m; i++ {
		distance := restrictions[i][0] - restrictions[i-1][0]
		// correct restriction for current building
		restrictions[i][1] = min(restrictions[i][1], restrictions[i-1][1]+distance)
	}

	// from right to left
	for i := m - 1; i >= 1; i-- {
		distance := restrictions[i][0] - restrictions[i-1][0]
		// correct restriction for current building
		restrictions[i][1] = min(restrictions[i][1], restrictions[i-1][1]+distance)
	}

	result := 0
	for i := 1; i < m; i++ {
		distance := restrictions[i][0] - restrictions[i-1][0]
		currMax := (distance + restrictions[i][1] + restrictions[i-1][1]) / 2

		result = max(result, currMax)
	}

	//fmt.Println(restrictions)
	return result
}
