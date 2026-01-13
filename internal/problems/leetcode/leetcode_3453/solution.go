package leetcode3453

import "math"

// https://leetcode.com/problems/separate-the-squares-in-a-grid/
//
// Time Complexity: O(n * log(maxY - minY)) ~> O(n * 50) where n is the number of squares
// Space Complexity: O(1)
func separateSquares(squares [][]int) float64 {
	var totalArea float64 = 0

	minY := math.Inf(1)
	maxY := math.Inf(-1)

	for _, square := range squares {
		y := float64(square[1])
		length := float64(square[2])

		minY = min(minY, y)
		maxY = max(maxY, y+length)

		totalArea += length * length
	}

	var halfArea float64 = totalArea / 2.0

	// use 50 iterations for better precision (enough for 1e-5)
	for range 50 {
		var midY float64 = minY + (maxY-minY)/2
		area := calculateArea(squares, midY, halfArea)
		if area < halfArea {
			minY = midY
		} else {
			maxY = midY
		}
	}

	return minY
}

func calculateArea(squares [][]int, y float64, limit float64) float64 {
	var below float64 = 0

	for _, square := range squares {
		y1 := float64(square[1])
		length := float64(square[2])
		y2 := y1 + length

		if y >= y2 {
			below += length * length
		} else if y > y1 {
			below += (y - y1) * length
		}

		if below > limit {
			return below
		}
	}
	return below
}
