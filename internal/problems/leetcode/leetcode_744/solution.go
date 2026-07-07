package leetcode744

// https://leetcode.com/problems/find-smallest-letter-greater-than-target/

// Solution1: Linear scan
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func nextGreatestLetter_solution1(letters []byte, target byte) byte {
	greater := letters[0]
	for _, letter := range letters {
		if letter > target {
			return letter
		}
	}
	return greater
}

// Solution2: Binary Search
//
// Time Complexity: O(log n)
// Space Complexity: O(1)
func nextGreatestLetter_solution2(letters []byte, target byte) byte {
	greater := binarySearch(letters, target)
	if greater == len(letters) {
		return letters[0]
	}
	return letters[greater]
}

func binarySearch(letters []byte, target byte) int {
	bad := -1
	good := len(letters)

	for good-bad > 1 {
		mid := bad + (good-bad)/2
		if letters[mid] > target {
			good = mid
		} else {
			bad = mid
		}
	}
	return good
}
