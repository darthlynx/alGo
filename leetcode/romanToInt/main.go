package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/roman-to-integer/
func main() {
	s := "MCMXCIV"
	i := romanToInt(s)
	fmt.Println(s, i)
}

func romanToInt(s string) int {
	chars := strings.Split(s, "")
	symbols := sympolsRepresentation()

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

func sympolsRepresentation() map[string]int {
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
