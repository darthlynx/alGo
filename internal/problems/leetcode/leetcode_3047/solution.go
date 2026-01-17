package leetcode3047

// https://leetcode.com/problems/find-the-largest-area-of-square-inside-two-rectangles/
//
// Time Complexity: O(n^2) where n is the number of rectangles
// Space Complexity: O(1)
func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)

	var maxArea int64 = 0
	for i := range n {
		for j := range n {
			if i == j {
				continue
			}
			height := min(topRight[i][1], topRight[j][1]) - max(bottomLeft[i][1], bottomLeft[j][1])
			width := min(topRight[i][0], topRight[j][0]) - max(bottomLeft[i][0], bottomLeft[j][0])

			// No overlap
			if height <= 0 || width <= 0 {
				continue
			}

			a := min(height, width)
			area := int64(a * a)
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}
