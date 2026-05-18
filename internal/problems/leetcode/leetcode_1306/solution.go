package leetcode1306

import "fmt"

// https://leetcode.com/problems/jump-game-iii/
//
// Time complexity: O(n).
// Space complexity: O(n).
func canReach(arr []int, start int) bool {
	q := queue{
		s:       []int{},
		headIdx: 0,
	}
	q.Add(start)

	for q.Size() > 0 {
		size := q.Size()
		for range size {
			idx, _ := q.Pop()
			if arr[idx] == 0 {
				return true
			}
			if arr[idx] < 0 { // already visited
				continue
			}

			if idx-arr[idx] >= 0 {
				q.Add(idx - arr[idx])
			}
			if idx+arr[idx] < len(arr) {
				q.Add(idx + arr[idx])
			}
			arr[idx] = -arr[idx] // mark as visited
		}
	}
	return false
}

type queue struct {
	s       []int
	headIdx int
}

func (q *queue) Add(idx int) {
	q.s = append(q.s, idx)
}

func (q *queue) Pop() (int, error) {
	if q.headIdx >= len(q.s) {
		return -1, fmt.Errorf("queue is empty")
	}
	num := q.s[q.headIdx]
	q.headIdx++
	return num, nil
}

func (q *queue) Size() int {
	return len(q.s) - q.headIdx
}
