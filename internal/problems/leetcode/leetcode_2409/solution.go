package leetcode2409

import (
	"strconv"
	"strings"
)

var months = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// https://leetcode.com/problems/count-days-spent-together/
//
// Time Complexity: O(1),
// Space Complexity: O(1).
func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	aliceArrive := toDays(arriveAlice)
	aliceLeave := toDays(leaveAlice)

	bobArrive := toDays(arriveBob)
	bobLeave := toDays(leaveBob)

	if bobArrive > aliceLeave || aliceArrive > bobLeave {
		return 0
	}

	return min(aliceLeave, bobLeave) - max(bobArrive, aliceArrive) + 1
}

func toDays(d string) int {
	parts := strings.Split(d, "-")
	day, _ := strconv.Atoi(parts[1])
	month, _ := strconv.Atoi(parts[0])

	days := 0
	for i := 0; i < month-1; i++ {
		days += months[i]
	}
	return days + day
}
