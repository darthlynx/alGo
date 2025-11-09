package leetcode13

import "strings"

// https://leetcode.com/problems/roman-to-integer/
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func romanToInt(s string) int {
	chars := strings.Split(s, "")
	symbols := symbolsRepresentation()

	var result, previous int

	for _, v := range chars {
		value := symbols[v]
		if value > previous {
			result += value - 2*previous
		} else {
			result += value
		}
		previous = value
	}
	return result
}

func symbolsRepresentation() map[string]int {
	sr := make(map[string]int)
	sr["I"] = 1
	sr["V"] = 5
	sr["X"] = 10
	sr["L"] = 50
	sr["C"] = 100
	sr["D"] = 500
	sr["M"] = 1000
	return sr
}
