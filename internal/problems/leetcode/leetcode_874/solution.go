package leetcode874

// https://leetcode.com/problems/walking-robot-simulation/
//
// Time Complexity: O(n + m), where n is the length of `commands` and m is the length of `obstacles`.
// Space Complexity: O(m), where m is the length of `obstacles`.
func robotSim(commands []int, obstacles [][]int) int {
	maxDistanceSquared := 0
	directions := [][]int{
		{0, 1},  // north
		{1, 0},  // east
		{0, -1}, // south
		{-1, 0}, // west
	}
	currentDirection := 0 // 0: north; 1: east; 2: south; 3: west

	// hash obstacles for faster search
	obs := make(map[position]struct{})
	for i := range obstacles {
		ob := toPosition(obstacles[i][0], obstacles[i][1])
		obs[ob] = struct{}{}
	}

	currentPos := position{
		x: 0,
		y: 0,
	}

	for i := 0; i < len(commands); i++ {
		command := commands[i]
		if command == -1 { // turn right
			currentDirection = (currentDirection + 1) % 4
			continue
		}
		if command == -2 { // turn left
			currentDirection = (currentDirection + 3) % 4
			continue
		}

		// move forward
		direction := directions[currentDirection]
		for step := 0; step < command; step++ {
			nextX := currentPos.x + direction[0]
			nextY := currentPos.y + direction[1]
			if _, exists := obs[toPosition(nextX, nextY)]; exists {
				break
			}
			currentPos.x = nextX
			currentPos.y = nextY
		}

		distance := currentPos.x*currentPos.x + currentPos.y*currentPos.y
		maxDistanceSquared = max(maxDistanceSquared, distance)
	}

	return maxDistanceSquared
}

type position struct {
	x int
	y int
}

func toPosition(xx, yy int) position {
	return position{
		x: xx,
		y: yy,
	}
}
