package leetcode2840

// https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-ii/
//
// Time Complexity: O(n) where n is the length of s1 and s2.
// Space Complexity: O(1) since the size of odd and even arrays is fixed at 26.
func checkStrings(s1 string, s2 string) bool {
	odd := make([]int, 26)
	even := make([]int, 26)

	for i := 0; i < len(s2); i++ {
		ch := int(s2[i] - 'a')
		if i%2 == 0 {
			even[ch]++
		} else {
			odd[ch]++
		}
	}

	for i := 0; i < len(s1); i++ {
		ch := int(s1[i] - 'a')
		if i%2 == 0 {
			if even[ch] <= 0 {
				return false
			}
			even[ch]--
		} else {
			if odd[ch] <= 0 {
				return false
			}
			odd[ch]--
		}
	}
	return true
}
