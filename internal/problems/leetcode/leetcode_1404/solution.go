package leetcode1404

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/
//
// Time Complexity: O(n) - where n is the length of the input string
// Space Complexity: O(1)
func numSteps(s string) int {
	steps := 0
	carry := 0 // carry from previous operation

	// iterate from the rightmost bit to the index 1
	// assuming that at index 0 we should leave 1 (our end result)
	// because `1` in binary is `1` in decimal, and we need to reduce it to 1
	for i := len(s) - 1; i >= 1; i-- {
		bit := int(s[i] - '0')

		x := bit + carry

		if x == 0 { // even, need to divide by 2 (right shift)
			steps++
		} else if x == 1 { // odd, +1 and then divide by 2
			steps += 2
			carry = 1
		} else { // x == 2, because bit == 1 and carry == 1
			steps++
			// do not need to add carry, it remains 1
		}
	}

	// in case carry remains 1, we end up in situation `10` which is equal to 2 in decimal
	// so to get the resulting 1 in decimal, we need to /2 (add one more operation)
	res := steps + carry
	return res
}
