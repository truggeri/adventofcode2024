package day14

import (
	"strconv"
	"strings"
)

const TIME = 100

type point struct {
	x, y int
}

type robot struct {
	position, velocity point
}

func Solve(input string, width, height int) uint {
	robots := parseInput(input)
	for i, r := range robots {
		robots[i] = moveRobot(r, TIME, width, height)
	}
	return safetyFactor(robots, width, height)
}

func parseInput(input string) []robot {
	robots := make([]robot, 0)
	for _, line := range strings.Split(input, "\n") {
		data := strings.Split(line, " ")
		if len(data) != 2 {
			continue
		}

		position := strings.Split(data[0][2:], ",")
		if len(position) != 2 {
			continue
		}
		x, _ := strconv.Atoi(position[0])
		y, _ := strconv.Atoi(position[1])
		p := point{x: x, y: y}

		velocity := strings.Split(data[1][2:], ",")
		if len(velocity) != 2 {
			continue
		}
		dx, _ := strconv.Atoi(velocity[0])
		dy, _ := strconv.Atoi(velocity[1])
		v := point{x: dx, y: dy}

		robots = append(robots, robot{position: p, velocity: v})
	}
	return robots
}

func moveRobot(r robot, time, width, height int) robot {
	newX := moveDirection(r.position.x, r.velocity.x, time, width)
	newY := moveDirection(r.position.y, r.velocity.y, time, height)
	p := point{x: newX, y: newY}
	r.position = p
	return r
}

func moveDirection(p, v, time, limit int) int {
	delta := abs(v*time) % limit
	if v > 0 {
		return (p + delta) % limit
	} else {
		return (p - delta + limit) % limit
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func safetyFactor(robots []robot, width, height int) uint {
	var quadrants [4][]robot
	halfX, halfY := width/2, height/2
	for _, r := range robots {
		if r.position.x == halfX || r.position.y == halfY {
			continue
		}

		if r.position.x < halfX {
			if r.position.y < halfY {
				quadrants[0] = append(quadrants[0], r)
			} else {
				quadrants[3] = append(quadrants[3], r)
			}
		} else {
			if r.position.y < halfY {
				quadrants[1] = append(quadrants[1], r)
			} else {
				quadrants[2] = append(quadrants[2], r)
			}
		}
	}

	return uint(len(quadrants[0]) * len(quadrants[1]) * len(quadrants[2]) * len(quadrants[3]))
}
