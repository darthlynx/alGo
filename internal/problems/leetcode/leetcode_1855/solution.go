package leetcode1855

// https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/
//
// Time complexity: O(m*log(n)).
// Space complexity: O(1).
func maxDistance(nums1 []int, nums2 []int) int {
	m := len(nums1)
	maxDist := 0
	for i := 0; i < m; i++ {
		j := binarySearch(nums2, nums1[i])
		if j >= i {
			maxDist = max(maxDist, j-i)
		}
	}
	return maxDist
}

func binarySearch(nums2 []int, num1 int) int {
	good := -1
	bad := len(nums2)
	for bad-good > 1 {
		mid := good + (bad-good)/2
		if nums2[mid] < num1 {
			bad = mid
		} else {
			good = mid
		}
	}
	return good
}
