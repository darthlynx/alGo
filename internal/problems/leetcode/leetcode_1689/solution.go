package leetcode1689

// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
//
// Time Complexity: O(n) - where n is the length of the input string
// Space Complexity: O(1)
//
// Explanation:
// minimum number of positive deci-binary numbers depends on the largest integer of n,
// for instance in "82734" 8 is largest so we should return 8
//  11111
//  11111
//  10111
//  10101
//  10100
//  10100
//  10100
//  10100
// +-----
//  82734
func minPartitions(n string) int {
	biggest := 0
	for i := 0; i < len(n); i++ {
		biggest = max(biggest, int(n[i]-'0'))
	}
	return biggest
}
