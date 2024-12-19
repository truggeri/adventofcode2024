package day10

import (
	"slices"
	"strconv"
	"strings"
)

const TRAIL_HEAD = 0
const MAX_HEIGHT = 9

type point struct {
	x, y int
}

func (p point) toString() string {
	return strings.Join([]string{strconv.Itoa(p.x), strconv.Itoa(p.y)}, ",")
}

type topMap [][]uint

func Solve(input string) uint {
	tm := parseInput(input)
	heads := findTrailHeads(tm)
	return walkTrails(tm, heads)
}

func parseInput(input string) [][]uint {
	result := make([][]uint, 0)
	for j, line := range strings.Split(input, "\n") {
		result = slices.Insert(result, j, make([]uint, 0))
		for i, h := range line {
			result[j] = slices.Insert(result[j], i, uint(h-'0'))
		}
	}
	return result
}

func findTrailHeads(tm topMap) []point {
	result := make([]point, 0)
	for j := 0; j < len(tm); j++ {
		for i := 0; i < len(tm[j]); i++ {
			if tm[j][i] == TRAIL_HEAD {
				result = append(result, point{x: i, y: j})
			}
		}
	}
	return result
}

func walkTrails(tm topMap, heads []point) uint {
	var result uint = 0
	for _, h := range heads {
		result += walkTrail(tm, h)
	}
	return result
}

func walkTrail(tm topMap, head point) uint {
	var result uint = 0
	paths := []point{head}
	seenPeaks := make(map[string]bool)
	for {
		if len(paths) == 0 {
			break
		}

		nextPaths := make([]point, 0)
		for _, p := range paths {
			if tm[p.y][p.x] == MAX_HEIGHT {
				if !seenPeaks[p.toString()] {
					seenPeaks[p.toString()] = true
					result++
				}
			} else {
				nextPaths = append(nextPaths, nextStepsFromPoint(tm, p)...)
			}
		}
		paths = nextPaths
	}
	return result
}

func nextStepsFromPoint(tm topMap, p point) []point {
	nextHeight := tm[p.y][p.x] + 1

	nextSteps := make([]point, 0)
	if p.y > 0 {
		if tm[p.y-1][p.x] == nextHeight {
			nextSteps = append(nextSteps, point{x: p.x, y: p.y - 1})
		}
	}
	if p.x > 0 {
		if tm[p.y][p.x-1] == nextHeight {
			nextSteps = append(nextSteps, point{x: p.x - 1, y: p.y})
		}
	}
	if p.y < len(tm)-1 {
		if tm[p.y+1][p.x] == nextHeight {
			nextSteps = append(nextSteps, point{x: p.x, y: p.y + 1})
		}
	}
	if p.x < len(tm[p.y])-1 {
		if tm[p.y][p.x+1] == nextHeight {
			nextSteps = append(nextSteps, point{x: p.x + 1, y: p.y})
		}
	}
	return nextSteps
}
