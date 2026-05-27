package leetcode3121

import "unicode"

// https://leetcode.com/problems/count-the-number-of-special-characters-ii/
//
// Time complexity: O(n),
// Space complexity: O(1)
func numberOfSpecialChars(word string) int {
	seen := make(map[rune][]int)
	special := 0

	for i, char := range word {
		if location, ok := seen[unicode.ToLower(char)]; ok {
			if unicode.IsUpper(char) {
				if location[1] >= 0 {
					continue
				} else {
					location[1] = i
				}
			} else {
				location[0] = i
			}
		} else {
			location := make([]int, 2)
			if unicode.IsUpper(char) {
				location[0] = -1
				location[1] = i
			} else {
				location[0] = i
				location[1] = -1
			}
			seen[unicode.ToLower(char)] = location
		}

	}

	for _, val := range seen {
		if val[0] >= 0 && val[1] >= 0 && val[1] > val[0] {
			special++
		}
	}

	return special
}
