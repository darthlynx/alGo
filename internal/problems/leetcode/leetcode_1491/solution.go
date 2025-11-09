package leetcode1491

import "math"

// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func average(salary []int) float64 {
	minSalary := math.MaxInt32
	maxSalary := math.MinInt32

	totalSum := 0

	for i := range salary {
		if salary[i] < minSalary {
			minSalary = salary[i]
		}
		if salary[i] > maxSalary {
			maxSalary = salary[i]
		}
		totalSum += salary[i]
	}

	// exclude min and max
	n := len(salary) - 2

	return float64(totalSum-minSalary-maxSalary) / float64(n)
}
