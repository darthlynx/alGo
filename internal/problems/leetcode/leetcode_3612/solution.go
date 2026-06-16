package leetcode3612

// https://leetcode.com/problems/process-string-with-special-operations-i/
//
// Time Complexity: O(n),
// Space Complexity: O(n).
func processStr(s string) string {
	stack := make([]byte, 0, len(s))

	for _, ch := range s {
		switch ch {
		case '*': // remove last
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case '#': // duplicate current result
			stack = append(stack, stack...)
		case '%': // reverse
			left := 0
			right := len(stack) - 1
			for left < right {
				stack[left], stack[right] = stack[right], stack[left]
				left++
				right--
			}
		default:
			stack = append(stack, byte(ch))
		}
	}

	return string(stack)
}
