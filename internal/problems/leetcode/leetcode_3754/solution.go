package leetcode3754

// https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-i/
//
// Time complexity: O(log(n))
// Space complexity: O(1)
func sumAndMultiply(n int) int64 {
	nonZero := 0
	for n > 0 {
		if n%10 != 0 {
			nonZero = nonZero*10 + n%10
		}
		n = n / 10
	}

	// reversing and finding sum of digits
	sum := 0
	x := 0
	for nonZero > 0 {
		sum += nonZero % 10
		x = x*10 + nonZero%10
		nonZero = nonZero / 10
	}

	return int64(x * sum)
}
