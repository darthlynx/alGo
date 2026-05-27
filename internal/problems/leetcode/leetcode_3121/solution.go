package leetcode3121

// https://leetcode.com/problems/count-the-number-of-special-characters-ii/
//
// Time complexity: O(n),
// Space complexity: O(1)
func numberOfSpecialChars(word string) int {
	firstUpper := [26]int{}
	lastLower := [26]int{}

	for i := range 26 {
		firstUpper[i] = -1
		lastLower[i] = -1
	}

	special := 0

	for i, char := range word {
		if char >= 'a' && char <= 'z' { // lower case
			lastLower[char-'a'] = i
		} else {
			if firstUpper[char-'A'] == -1 {
				firstUpper[char-'A'] = i
			}
		}
	}

	for i := range 26 {
		if lastLower[i] >= 0 && firstUpper[i] >= 0 && lastLower[i] < firstUpper[i] {
			special++
		}
	}

	return special
}
