package leetcode3653

// https://leetcode.com/problems/xor-after-range-multiplication-queries-i/
//
// Time Complexity: O(n + m*n), where n is the length of `nums` and m is the length of `queries`.
// Space Complexity: O(1).
func xorAfterQueries(nums []int, queries [][]int) int {
	const modulo = 1_000_000_007
	for _, query := range queries {
		for i := query[0]; i <= query[1]; i += query[2] {
			nums[i] = (nums[i] * query[3]) % modulo
		}
	}

	xor := 0
	for i := range nums {
		xor = xor ^ nums[i]
	}
	return xor
}
