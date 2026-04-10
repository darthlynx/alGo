package leetcode3740

import "math"

// https://leetcode.com/problems/minimum-distance-between-three-equal-elements-i/
//
// Time complexity: O(n).
// Space complexity: O(n).
func minimumDistance(nums []int) int {
	m := make(map[int][]int)
	for i, num := range nums {
		m[num] = append(m[num], i)
	}

	minDistance := math.MaxInt
	for _, v := range m {
		if len(v) < 3 {
			continue
		}
		for i := 0; i < len(v)-2; i++ {
			current := abs(v[i]-v[i+1]) + abs(v[i+1]-v[i+2]) + abs(v[i+2]-v[i])
			minDistance = min(minDistance, current)
		}
	}

	if minDistance == math.MaxInt {
		return -1
	}
	return minDistance
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
