package leetcode2553

// https://leetcode.com/problems/separate-the-digits-in-an-array/
//
// Time complexity: O(n*m) where n is the length of nums and m is the number of digits in the largest number.
// Space complexity: O(n*m) for the result array.
func separateDigits(nums []int) []int {
	var result []int
	for _, num := range nums {
		result = append(result, getDigits(num)...)
	}
	return result
}

func getDigits(num int) []int {
	var digits []int
	for num > 0 {
		digit := num % 10
		digits = append(digits, digit)
		num = num / 10
	}
	reverse(digits)
	return digits
}

func reverse(digits []int) {
	left := 0
	right := len(digits) - 1
	for left < right {
		digits[left], digits[right] = digits[right], digits[left]
		left++
		right--
	}
}
