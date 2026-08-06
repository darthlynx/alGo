package leetcode3345

// https://leetcode.com/problems/smallest-divisible-digit-product-i/
//
// Time Complexity: O(log(n)), where n is the value of the input number
// Space Complexity: O(1)
func smallestNumber(n int, t int) int {
	// any number ending in 0 has product 0, divisible by any t, so this always terminates
	for getDigitsProduct(n)%t != 0 {
		n++
	}
	return n
}

func getDigitsProduct(n int) int {
	res := 1
	for n > 0 {
		res *= n % 10
		n /= 10
	}
	return res
}
