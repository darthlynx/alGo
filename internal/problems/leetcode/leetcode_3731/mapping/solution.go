package mapping

import "math"

// https://leetcode.com/problems/find-missing-elements/
//
// Time Complexity: O(n + m), where m is the range between smallest and biggest numbers
// Space Complexity: O(n + m)
func findMissingElements(nums []int) []int {
	seen := make(map[int]struct{})
	smallest := math.MaxInt
	biggest := math.MinInt
	for _, num := range nums {
		seen[num] = struct{}{}
		smallest = min(smallest, num)
		biggest = max(biggest, num)
	}

	var res []int
	for i := smallest; i < biggest; i++ {
		if _, ok := seen[i]; !ok {
			res = append(res, i)
		}
	}

	return res
}
