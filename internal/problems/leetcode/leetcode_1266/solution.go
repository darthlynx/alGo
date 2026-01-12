package leetcode1266

// https://leetcode.com/problems/minimum-time-visiting-all-points/
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func minTimeToVisitAllPoints(points [][]int) int {
	time := 0

	for i := 0; i < len(points)-1; i++ {
		dx := abs(points[i+1][0] - points[i][0])
		dy := abs(points[i+1][1] - points[i][1])

		// min(dx, dy) + (max(dx, dy) - min(dx, dy)) = max(dx, dy)
		time += max(dx, dy)
	}
	return time
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
