package leetcode3

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Solution1: Using map to store characters
func lengthOfLongestSubstringMap(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	var max int
	charset := make(map[byte]struct{})
	var start int
	charset[s[start]] = struct{}{}
	index := 1

	for index < len(s) {
		v := s[index]
		if _, ok := charset[v]; ok {
			max = maxNum(max, len(charset))
			charset = make(map[byte]struct{})
			start++
			index = start + 1
			charset[s[start]] = struct{}{}
		} else {
			charset[v] = struct{}{}
			max = maxNum(max, len(charset))
			index++
		}
	}
	return max
}

// Solution2: Using sliding window with array
func lengthOfLongestSubstringWindow(s string) int {
	var max int
	presence := make([]int, 255) // will contain unicode characters by their bytes
	for i := range presence {
		presence[i] = -1
	}

	lastRepeated := -1
	for i := 0; i < len(s); i++ {
		lastRepeated = maxNum(lastRepeated, presence[s[i]])
		// find the difference between current position and last repeated value
		// (which contains previous position of this character)
		// it will give us current length of substring without repeated characters
		max = maxNum(max, i-lastRepeated)
		presence[s[i]] = i
	}
	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
