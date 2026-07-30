package leetcode3014

// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-i/
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func minimumPushes(word string) int {
	const numOfKeys = 8
	result := 0
	for i := 0; i < len(word); i++ {
		result += i/numOfKeys + 1
	}
	return result
}
