package leetcode3622

// https://leetcode.com/problems/check-divisibility-by-digit-sum-and-product/
//
// Time Complexity: O(log n), where n is the input number
// Space Complexity: O(1)
func checkDivisibility(n int) bool {
	sum := 0
	product := 1
	m := n
	for m > 0 {
		digit := m % 10
		sum += digit
		product *= digit
		m /= 10
	}
	return n%(sum+product) == 0
}
