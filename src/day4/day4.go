package day4

import (
	"slices"
	"strings"
)

type strategy func([][]rune, int, int) bool

func Solve(input string) uint {
	puzzle := parseInput(input)
	ss := []strategy{builder(0, -1), builder(0, 1), builder(-1, 0), builder(1, 0), builder(-1, -1), builder(1, -1), builder(-1, 1), builder(1, 1)}
	var result uint = 0
	for j, line := range puzzle {
		for i := range line {
			for _, s := range ss {
				if s(puzzle, i, j) {
					result++
				}
			}
		}
	}
	return result
}

func parseInput(input string) [][]rune {
	result := make([][]rune, 0)
	for j, line := range strings.Split(input, "\n") {
		result = slices.Insert(result, j, make([]rune, 0))
		for _, r := range line {
			result[j] = append(result[j], r)
		}
	}
	return result
}

func builder(right, up int) strategy {
	word := [4]rune{'X', 'M', 'A', 'S'}
	length := len(word)
	return func(puzzle [][]rune, x, y int) bool {
		if (up < 0 && y < length-1) || (up > 0 && y > len(puzzle)-length) {
			return false
		}
		if (right < 0 && x < length-1) || (right > 0 && x > len(puzzle)-length) {
			return false
		}

		for i, r := range word {
			if puzzle[y+(i*up)][x+(i*right)] != r {
				return false
			}
		}
		return true
	}
}
