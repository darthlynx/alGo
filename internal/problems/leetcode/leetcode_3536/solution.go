package leetcode3536

// https://leetcode.com/problems/maximum-product-of-two-digits/
//
// Time Complexity: O(log(n))
// Space Complexity: O(1)
func maxProduct(n int) int {
	max1 := 0
	max2 := 0

	for n > 0 {
		digit := n % 10
		if digit >= max1 {
			max2 = max1
			max1 = digit
		} else if digit >= max2 {
			max2 = digit
		}
		n = n / 10
	}

	return max1 * max2
}
