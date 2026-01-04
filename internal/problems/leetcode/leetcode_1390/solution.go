package leetcode1390

// https://leetcode.com/problems/sum-of-four-divisors/
//
// Time Complexity: O(n * sqrt(m)), where n is the length of nums and m is the maximum number in nums
// Space Complexity: O(1)
func sumFourDivisors(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += getDivisorsSum(num)
	}
	return sum
}

func getDivisorsSum(num int) int {
	count := 0
	sum := 0
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			count++
			if i != num/i {
				sum += num / i
				count++
			}
			if count > 4 {
				return 0
			}
		}
	}
	if count == 4 {
		return sum
	}
	return 0
}
