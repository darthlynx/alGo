package leetcode3612

// https://leetcode.com/problems/process-string-with-special-operations-i/
//
// Time Complexity: O(n + total length processed by '#' and '%' operations).
// In the worst case, this can be O(2^n) due to repeated duplication.
// Space Complexity: O(m), where m is the maximum generated string length.
func processStr(s string) string {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '*': // remove last
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case '#': // duplicate current result
			stack = append(stack, stack...)
		case '%':
			reverse(stack)
		default:
			stack = append(stack, byte(s[i]))
		}
	}

	return string(stack)
}

func reverse(chars []byte) {
	left := 0
	right := len(chars) - 1
	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
}
