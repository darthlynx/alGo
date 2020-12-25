package main

import (
	"fmt"
	"strings"
)

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
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

type void struct{}

func lengthOfLongestSubstring(s string) int {
	input := strings.Split(s, "")
	var max int
	charset := make(map[string]void)
	if len(input) <= 1 {
		return len(input)
	}
	var start int
	charset[input[start]] = void{}
	index := 1
	for index < len(input) {
		v := input[index]
		// fmt.Printf("start: %v\t index: %v\t charset: %v\t cur: %v\t prev:%v\n", start, index, charset, v, input[index-1])
		if _, ok := charset[v]; ok {
			max = maxNum(max, len(charset))
			charset = make(map[string]void)
			start++
			index = start + 1
			charset[input[start]] = void{}
		} else {
			charset[v] = void{}
			max = maxNum(max, len(charset))
			index++
		}
	}
	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
