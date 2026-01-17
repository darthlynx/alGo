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

            // not overlapping by X
            if bottomLeft[i][0] >= topRight[j][0] || bottomLeft[j][0] >= topRight[i][0] {
                continue
            }
            // not overlapping by Y
            if topRight[i][1] <= bottomLeft[j][1] || topRight[j][1] <= bottomLeft[i][1] {
                continue
            }

            yTop := min(topRight[i][1], topRight[j][1])
            yBottom := max(bottomLeft[i][1], bottomLeft[j][1])
            yy := yTop - yBottom

            xRight := min(topRight[i][0], topRight[j][0])
            xLeft := max(bottomLeft[i][0], bottomLeft[j][0])
            xx := xRight - xLeft

            a := min(xx, yy)
            area := int64(a * a)
            maxArea = max(maxArea, area)
        }
    }

    return maxArea
}
