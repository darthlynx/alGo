package leetcode1390

// https://leetcode.com/problems/sum-of-four-divisors/
//
// Time Complexity: O(n * sqrt(m)), where n is the length of nums and m is the maximum number in nums
// Space Complexity: O(1)
func sumFourDivisors(nums []int) int {
	sum := 0
	for _, num := range nums {
		divisorsSum := getDivisorsSum(num)
		sum += divisorsSum
	}

	return sum
}

func getDivisorsSum(num int) int {
	divisors := []int{}

	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
			if i != num/i {
				divisors = append(divisors, num/i)
			}

			if len(divisors) > 4 {
				return 0
			}
		}
	}

	if len(divisors) == 4 {
		return divisors[0] + divisors[1] + divisors[2] + divisors[3]
	}

	return 0
}
