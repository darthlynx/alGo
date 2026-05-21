package leetcode3043

// https://leetcode.com/problems/find-the-length-of-the-longest-common-prefix/
//
// Time complexity: O((m + n) * d), where m = len(arr1), n = len(arr2), d = max number of digits in any number in arr1 or arr2
// Space complexity: O(m * d), where m = len(arr1), d = max number of digits in any number in arr1
func longestCommonPrefix(arr1 []int, arr2 []int) int {
	set := make(map[int]struct{}, 10*len(arr1))

	// collect all prefixes for nums in arr1
	for _, num := range arr1 {
		for num > 0 {
			set[num] = struct{}{}
			num = num / 10
		}
	}

	maxPrefixLen := 0
	for _, num := range arr2 {
		for num > 0 {
			if _, ok := set[num]; ok {
				maxPrefixLen = max(maxPrefixLen, getLength(num))
				break // no need to check shorter prefixes for this number
			}
			num = num / 10
		}
	}
	return maxPrefixLen
}

func getLength(num int) int {
	length := 0
	for num > 0 {
		length++
		num = num / 10
	}
	return length
}
