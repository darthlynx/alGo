package leetcode9

// https://leetcode.com/problems/palindrome-number/description/
//
// Time Complexity: O(log n) - where n is the input number
// Space Complexity: O(1)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev := reverse(x)
	return rev == x
}

func reverse(x int) int {
	r := 0
	for x > 0 {
		r = r*10 + x%10
		x = x / 10
	}
	return r
}
