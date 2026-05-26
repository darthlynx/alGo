package leetcode3120

import (
	"unicode"
)

// https://leetcode.com/problems/count-the-number-of-special-characters-i/
//
// Time complexity: O(n),
// Space complexity: O(1) - at most 52 characters (26 lowercase + 26 uppercase)
func numberOfSpecialChars(word string) int {
    contains := make(map[rune]bool, 2*26)
    count := 0

    for _, char := range word {
        if contains[unicode.ToUpper(char)] && contains[unicode.ToLower(char)] {
            continue
        }
        if unicode.IsLower(char) && contains[unicode.ToUpper(char)] {
            count++
        } else if unicode.IsUpper(char) && contains[unicode.ToLower(char)] {
            count++
        }
        contains[char] = true
    }

    return count
}
