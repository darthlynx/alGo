package leetcode3318

import "sort"

// https://leetcode.com/problems/sum-of-xor-beauty-of-all-subarrays-of-size-k/
// Time Complexity: O(n * k log k)
// Space Complexity: O(k)
func findXSum(nums []int, k int, x int) []int {
	m := len(nums) - k + 1
	result := make([]int, m)
	for i := 0; i < m; i++ {
		result[i] = xSum(nums, i, k, x)
	}
	return result
}

func xSum(nums []int, start, k, x int) int {
	freq := make(map[int]int)
	for i := start; i < start+k; i++ {
		freq[nums[i]]++
	}

	entries := make([][2]int, 0, len(freq))
	for num, count := range freq {
		entries = append(entries, [2]int{num, count})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i][1] == entries[j][1] {
			return entries[i][0] > entries[j][0]
		}
		return entries[i][1] > entries[j][1]
	})

	sum := 0
	for i := 0; i < x && i < len(entries); i++ {
		sum += entries[i][0] * entries[i][1]
	}
	return sum
}
