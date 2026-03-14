package leetcode1415

import "slices"

// https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/
//
// Time complexity: O(2^n * n) in practice, due to generation plus final sorting.
// Space complexity: O(2^n * n) for storing generated strings.
func getHappyString(n int, k int) string {
    happyStrings := make(map[string]struct{}, k)
    generate("", n, k, happyStrings)
    if len(happyStrings) < k {
        return ""
    }

    sorted := []string{}
    for key := range happyStrings {
        sorted = append(sorted, key)
    }
    slices.Sort(sorted)

    return sorted[k-1]
}

func generate(base string, n int, k int, happyStrings map[string]struct{}) {
    if len(happyStrings) == k {
        return
    }
    if len(base) == n {
        happyStrings[base] = struct{}{}
        return
    }

    if !endsWith(base, 'a') {
        generate(base+"a", n, k, happyStrings)
    }
    if !endsWith(base, 'b') {
        generate(base+"b", n, k, happyStrings)
    }
    if !endsWith(base, 'c') {
        generate(base+"c", n, k, happyStrings)
    }
}

func endsWith(s string, suffix byte) bool {
    if len(s) == 0 {
        return false
    }
    return s[len(s)-1] == suffix
}
