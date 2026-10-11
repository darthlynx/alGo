package leetcode921

// https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
//
// Time complexity: O(n)
// Space complexity: O(1)
func minAddToMakeValid(s string) int {
	reqOpen := 0
	reqClosed := 0
	for _, ch := range s {
		if ch == '(' {
			reqOpen++
		} else {
			if reqOpen == 0 {
				reqClosed++
			} else {
				reqOpen--
			}
		}
	}
	return reqOpen + reqClosed
}
