package leetcode1344

import "math"

// https://leetcode.com/problems/angle-between-hands-of-a-clock/
//
// Time Complexity: O(1).
// Space Complexity: O(1).
func angleClock(hour int, minutes int) float64 {
	degreesPerHour := 360 / 12   // 30 degrees in 1 hour
	degreesPerMinute := 360 / 60 // 6 degrees in 1 minute

	if hour == 12 { // to avoid 360 deg overflow
		hour = 0
	}

	// how to calculate extra angle for hour hand:
	// 30deg = 60 minutes
	// x deg = 15 minute
	// x = 30 * 15 / 60
	roughHourAngle := float64(hour * degreesPerHour)
	extraHourAngle := float64(minutes*degreesPerHour) / 60
	hourAngle := roughHourAngle + extraHourAngle

	minuteAngle := float64(minutes * degreesPerMinute)

	diff := math.Abs(minuteAngle - hourAngle)
	if diff > 180 {
		diff = 360 - diff
	}
	return diff
}
