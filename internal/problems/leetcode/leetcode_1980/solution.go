package leetcode1980

// https://leetcode.com/problems/find-unique-binary-string/
//
// Time Complexity: O(2^n) - where n is the length of the input array
// Space Complexity: O(n) - for the recursive call stack and the set
func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	var notSeen string
	set := make(map[string]bool)
	for _, num := range nums {
		set[num] = true
	}
	generate("", set, n, &notSeen)

	return notSeen
}

func generate(base string, set map[string]bool, n int, notSeen *string) {
	if *notSeen != "" {
		return
	}
	if len(base) == n {
		if !set[base] {
			*notSeen = base
		}
		return
	}

	generate(base+"0", set, n, notSeen)
	generate(base+"1", set, n, notSeen)
}
