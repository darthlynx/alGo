package leetcode2540

// https://leetcode.com/problems/minimum-common-value/
//
// Time complexity: O(n + m).
// Space complexity: O(1).
func getCommon(nums1 []int, nums2 []int) int {
	i1 := 0
	i2 := 0

	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] == nums2[i2] {
			return nums1[i1]
		}

		if nums1[i1] < nums2[i2] {
			i1++
		} else {
			i2++
		}
	}

	return -1
}
