package day6

import (
	"slices"
	"strings"
)

type direction uint

const (
	Up direction = iota
	Right
	Down
	Left
)

type room struct {
	layout [][]square
	guard  guard
}

type square struct {
	visited  bool
	obstacle bool
}

type guard struct {
	direction direction
	location
}

type location struct {
	x int
	y int
}

func Solve(input string) uint {
	rm := parseInput(input)
	return solve(rm)
}

func parseInput(input string) room {
	var rm room
	for j, l := range strings.Split(input, "\n") {
		rm.layout = slices.Insert(rm.layout, j, make([]square, 0))
		for i, ch := range l {
			sq := square{}
			if ch == '#' {
				sq.obstacle = true
			} else if ch != '.' {
				rm.guard = guard{
					direction: parseDirection(ch),
					location: location{
						x: i,
						y: j,
					},
				}
				sq.visited = true
			}
			rm.layout[j] = slices.Insert(rm.layout[j], i, sq)
		}
	}
	return rm
}

func parseDirection(input rune) direction {
	switch input {
	case '>':
		return Right
	case '<':
		return Left
	case 'v':
		return Down
	}
	return Up
}

func solve(rm room) uint {
	var moves, unique uint = 0, 1
	nextLocation := rm.guard.location
	for {
		moves++
		x, y := adjustLocation(rm.guard.direction)
		nextLocation.x += x
		nextLocation.y += y

		if offGrid(nextLocation, rm) {
			break
		}

		nextSquare := rm.layout[nextLocation.y][nextLocation.x]
		if nextSquare.obstacle {
			rm.guard.direction += 1
			if rm.guard.direction > Left {
				rm.guard.direction = Up
			}
			nextLocation = rm.guard.location
		} else {
			if !nextSquare.visited {
				unique++
				rm.layout[nextLocation.y][nextLocation.x].visited = true
			}
			rm.guard.location = nextLocation
		}
	}
	return unique
}

func adjustLocation(input direction) (int, int) {
	switch input {
	case Right:
		return 1, 0
	case Left:
		return -1, 0
	case Down:
		return 0, 1
	case Up:
		return 0, -1
	}
	return 0, 0
}

func offGrid(loc location, rm room) bool {
	return loc.x < 0 || loc.y < 0 || loc.x >= len(rm.layout) || loc.y >= len(rm.layout[0])
}
