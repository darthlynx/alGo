package leetcode1306

import "fmt"

// https://leetcode.com/problems/jump-game-iii/
//
// Time complexity: O(n).
// Space complexity: O(n).
func canReach(arr []int, start int) bool {
	q := queue{
		s: []int{},
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
	s []int
}

func (q *queue) Add(idx int) {
	q.s = append(q.s, idx)
}

func (q *queue) Pop() (int, error) {
	if len(q.s) == 0 {
		return -1, fmt.Errorf("queue is empty")
	}
	num := q.s[0]
	q.s = q.s[1:]
	return num, nil
}

func (q *queue) Size() int {
	return len(q.s)
}
