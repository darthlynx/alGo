package leetcode2778

// https://leetcode.com/problems/sum-of-squares-of-special-elements/
//
// Time complexity: O(n)
// Space complexity: O(1)
func sumOfSquares(nums []int) int {
	n := len(nums)
	sum := 0
	for i, val := range nums {
		if n%(i+1) == 0 {
			sum += val * val
		}
	}
	return sum
}
