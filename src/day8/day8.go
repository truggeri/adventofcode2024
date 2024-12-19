package day8

import (
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func (p point) toString() string {
	return strings.Join([]string{strconv.Itoa(p.x), strconv.Itoa(p.y)}, ",")
}

func (p point) onMap() bool {
	return p.x >= 0 && p.x < xMax && p.y >= 0 && p.y < yMax
}

var xMax, yMax = 0, 0

func Solve(input string) uint {
	return solve(parseInput(input))
}

func parseInput(input string) map[rune][]point {
	xMax, yMax = 0, len(strings.Split(input, "\n"))
	result := make(map[rune][]point)
	for j, line := range strings.Split(input, "\n") {
		if xMax == 0 {
			xMax = len(line)
		}
		for i, chr := range line {
			if chr == '.' {
				continue
			}

			pt := point{x: i, y: j}
			result[chr] = append(result[chr], pt)
		}
	}
	return result
}

func solve(input map[rune][]point) uint {
	antinodes := make(map[string]interface{})
	for _, locations := range input {
		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				var dx int = locations[i].x - locations[j].x
				var dy int = locations[i].y - locations[j].y
				antiNodeNearI := point{x: locations[i].x + dx, y: locations[i].y + dy}
				antiNodeNearJ := point{x: locations[j].x - dx, y: locations[j].y - dy}

				if antiNodeNearI.onMap() {
					antinodes[antiNodeNearI.toString()] = true
				}
				if antiNodeNearJ.onMap() {
					antinodes[antiNodeNearJ.toString()] = true
				}
			}
		}
	}
	return uint(len(antinodes))
}
