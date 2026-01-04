package leetcode1390

// https://leetcode.com/problems/sum-of-four-divisors/
//
// Time Complexity: O(n * sqrt(m)) where n is the length of nums and m is the max number in nums
// Space Complexity: O(n)
func sumFourDivisors(nums []int) int {
	divisors := make(map[int]map[int]struct{})

	for i := range nums {
		if divisors[nums[i]] != nil {
			continue
		} else {
			divisors[nums[i]] = getDivisors(nums[i])
		}
	}

	count := 0
	for _, num := range nums {
		divs := divisors[num]
		if len(divs) == 4 {
			for key := range divs {
				count += key
			}
		}
	}
	return count
}

func getDivisors(num int) map[int]struct{} {
	divisors := make(map[int]struct{})

	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			divisors[i] = struct{}{}
			divisors[num/i] = struct{}{}
		}
	}
	return divisors
}
