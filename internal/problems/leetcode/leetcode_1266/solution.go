package leetcode1266

// https://leetcode.com/problems/minimum-time-visiting-all-points/
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func minTimeToVisitAllPoints(points [][]int) int {
	time := 0
	for i := 0; i < len(points)-1; i++ {
		x := points[i][0]
		y := points[i][1]

		xx := points[i+1][0]
		yy := points[i+1][1]

		dx := abs(xx - x)
		dy := abs(yy - y)

		time += min(dx, dy)
		time += max(dx, dy) - min(dx, dy)
	}
	return time
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
