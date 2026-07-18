package euclideanalgo

// https://leetcode.com/problems/find-greatest-common-divisor-of-array/
//
// Time Complexity: O(n + log mn)
// Space Complexity: O(1)
func findGCD(nums []int) int {
	mn := nums[0]
	mx := nums[0]

	for _, num := range nums {
		mn = min(mn, num)
		mx = max(mx, num)
	}

	return gcd(mn, mx)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
