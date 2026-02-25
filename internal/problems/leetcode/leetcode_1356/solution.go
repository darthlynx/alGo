package leetcode1356

import (
	"math/bits"
	"sort"
)

// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/
//
// Time Complexity: O(n + k log k) - where k is the number of elements in each bucket
// Space Complexity: O(n) - due to buckets and result array
func sortByBits(arr []int) []int {
	// Implementing Bucket sort here

    // max number of buckets is 15
    // each bucket has from 0 to 14 '1'-s
    // because arr[i] <= 10_000, and it needs 14 bits to represent
    buckets := make([][]int, 15)

    // put each number into the right bucket
    for _, num := range arr {
        bCount := bits.OnesCount(uint(num))
        buckets[bCount] = append(buckets[bCount], num)
    }

    res := make([]int, len(arr))
    idx := 0
    // concat the numbers from each bucket to result
    for i := range buckets {
        length := len(buckets[i])
        if length == 0 { // skip if bucket is empty
            continue
        }

        // sort nums in the bucket and append them to the result
        sort.Ints(buckets[i])
        for j := range buckets[i] {
            res[idx] = buckets[i][j]
            idx++
        }
    }

    return res
}
