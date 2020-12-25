package main

import (
	"fmt"
	"strings"
)

func main() {
	m := map[string]int{}
	m["pwwkew"] = 3
	m["abcabcbb"] = 3
	m["bbb"] = 1
	m[""] = 0
	m["abc"] = 3
	m["au"] = 2

	for k, v := range m {
		actual := lengthOfLongestSubstring(k)
		fmt.Println(k, v, actual, actual == v)
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
	var index int
	for index < len(input) {
		v := input[index]
		//fmt.Println("value", v)
		if _, ok := charset[v]; ok {
			max = maxNum(max, len(charset))
			start++
			index = start + 1
			charset = make(map[string]void)
		} else {
			max = maxNum(max, len(charset))
			index++
		}
		charset[v] = void{}
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
