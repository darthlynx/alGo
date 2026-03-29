package leetcode2839

// https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-i/
//
// Time Complexity: O(1) since the length of s1 and s2 is fixed at 4.
// Space Complexity: O(1).
func canBeEqual(s1 string, s2 string) bool {
	chars := []byte(s2)
	for i := range 4 {
		if s1 == string(chars) {
			return true
		}
		if i%2 == 0 {
			chars[0], chars[2] = chars[2], chars[0]
		} else {
			chars[1], chars[3] = chars[3], chars[1]
		}
	}
	return false
}

// Time Complexity: O(1) since the length of s1 and s2 is fixed at 4.
// Space Complexity: O(1).
func canBeEqual_2(s1 string, s2 string) bool {
	same := func(a, b, c, d byte) bool {
		return (a == c && b == d) || (a == d && b == c)
	}
	return same(s1[0], s1[2], s2[0], s2[2]) && same(s1[1], s1[3], s2[1], s2[3])
}
