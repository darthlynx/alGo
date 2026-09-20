package leetcode_3498

// https://leetcode.com/problems/reverse-degree-of-a-string/
//
// Time complexity: O(n)
// Space complexity: O(1)
func reverseDegree(s string) int {
	sum := 0

	for i, v := range s {
		val := 26 - int(v-'a')
		sum += val * (i + 1)
	}
	return sum
}
