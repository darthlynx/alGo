package leetcode3517

// https://leetcode.com/problems/smallest-palindromic-rearrangement-i/
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func smallestPalindrome(s string) string {
	n := len(s)
	chCount := [26]int{}
	for _, ch := range s {
		chCount[ch-'a']++
	}

	result := make([]byte, n)
	left := 0
	right := n - 1

	for i := range chCount {
		ch := byte(i + 'a')
		for chCount[i] >= 2 {
			result[left] = ch
			result[right] = ch
			chCount[i] -= 2
			left++
			right--
		}
		// uneven number of chars (could be only one in palindrome, so we put it in the middle)
		if chCount[i] > 0 {
			result[n/2] = ch
		}
	}

	return string(result)
}
