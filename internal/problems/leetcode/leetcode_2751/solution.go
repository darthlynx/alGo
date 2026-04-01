package leetcode2751

import "slices"

// https://leetcode.com/problems/robot-collisions/
//
// Time Complexity: O(n log n) where n is the number of robots due to sorting.
// Space Complexity: O(n) for the sorted array and the stack.
func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)

	sorted := make([]robot, 0, n)
	for i := 0; i < n; i++ {
		sorted = append(sorted, robot{
			position:  positions[i],
			health:    healths[i],
			direction: string(directions[i]),
			originIdx: i,
		})
	}
	// sort robots by position
	slices.SortFunc(sorted, func(a, b robot) int {
		if a.position < b.position {
			return -1
		}
		if a.position > b.position {
			return 1
		}
		return 0
	})

	s := Stack{
		robots: []robot{},
	}

	for i := 0; i < len(sorted); i++ {
		curr := sorted[i]

		alive := true

		for alive { // check for collisions
			prev, exists := s.Peek()
			if !exists || !(prev.direction == "R" && curr.direction == "L") {
				break
			}
			s.Pop()

			if prev.health < curr.health { // current is stronger
				curr.health--
			} else if prev.health == curr.health { // both are equally strong, eliminate both
				alive = false
			} else { // previous is stronger
				prev.health--
				s.Push(prev)
				alive = false
			}
		}
		if alive && curr.health > 0 {
			s.Push(curr)
		}
	}

	// to preserve the result in the original order
	resRobots := s.robots
	slices.SortFunc(resRobots, func(a, b robot) int {
		if a.originIdx < b.originIdx {
			return -1
		}
		if a.originIdx > b.originIdx {
			return 1
		}
		return 0
	})
	result := make([]int, 0, len(resRobots))
	for i := 0; i < len(s.robots); i++ {
		result = append(result, resRobots[i].health)
	}

	return result
}

type robot struct {
	position  int
	health    int
	direction string
	originIdx int
}

type Stack struct {
	robots []robot
}

func (s *Stack) Push(r robot) {
	s.robots = append(s.robots, r)
}

func (s *Stack) Peek() (robot, bool) {
	if len(s.robots) == 0 {
		return robot{}, false
	}
	return s.robots[len(s.robots)-1], true
}

func (s *Stack) Pop() (robot, bool) {
	if len(s.robots) == 0 {
		return robot{}, false
	}
	r := s.robots[len(s.robots)-1]
	s.robots = s.robots[:len(s.robots)-1]
	return r, true
}
