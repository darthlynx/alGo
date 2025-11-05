package leetcode1

// https://leetcode.com/problems/two-sum/
//
// Time Complexity: O(n), Space Complexity: O(n)
func TwoSum(nums []int, target int) []int {
	checked := make(map[int]int)

	for i, el := range nums {
		if v, ok := checked[target-el]; ok {
			return []int{i, v}
		}
		checked[el] = i
	}
	return []int{-1, -1}
}
