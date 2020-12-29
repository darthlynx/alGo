package main

import (
	"fmt"
)

func main() {
	m := map[string]int{}
	m["pwwkew"] = 3
	m["abcabcbb"] = 3
	m["bbb"] = 1
	m[""] = 0
	m["abc"] = 3
	m["au"] = 2
	m["dvdf"] = 3
	m["ckilbkd"] = 5

	for k, v := range m {
		actual := lengthOfLongestSubstring(k)
		fmt.Println(k, "\t", v, actual, "\t", actual == v)
	}
}

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(str string) int {
	var max int
	presence := getEmptySlice(255) // will contain unicode characters by their bytes
	lastRepeated := -1
	for i := 0; i<len(str); i++ {
		lastRepeated = maxOf(lastRepeated, presence[str[i]])
		max = maxOf(max, i - lastRepeated)
		presence[str[i]] = i
	}
	return max
}

func getEmptySlice(size int) []int {
	var sl []int
	for i := 0; i < size; i++ {
		sl = append(sl, -1)
	}
	return sl
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
