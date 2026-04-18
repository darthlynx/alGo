package leetcode3783

// https://leetcode.com/problems/mirror-distance-of-an-integer/
//
// Time complexity: O(log(n)).
// Space complexity: O(1).
func mirrorDistance(n int) int {
	reversed := reverse(n)
	return abs(n - reversed)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func reverse(n int) int {
	res := 0
	for n > 0 {
		digit := n % 10
		res = res*10 + digit
		n = n / 10
	}
	return res
}
