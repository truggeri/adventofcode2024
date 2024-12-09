package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const DELTA_MAX = 3
const DELTA_MIN = 1

type level uint
type report []level

func main() {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	solution := Solve(input)
	fmt.Println("Solution:", solution)
}

func Solve(input string) uint {
	return solve(parseInput(input))
}

func parseInput(input string) []report {
	var result []report
	reports := strings.Split(input, "\n")
	for i, report := range reports {
		result = slices.Insert(result, i, []level{})
		levels := strings.Split(report, " ")
		for j, num := range levels {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic("Cannot parse input")
			}
			result[i] = slices.Insert(result[i], j, level(n))
		}
	}
	return result
}

func solve(input []report) uint {
	var result uint = 0
	for _, report := range input {
		if isValidReport(report) {
			result += 1
		}
	}
	return result
}

func isValidReport(report report) bool {
	if len(report) < 2 {
		return true
	}

	prev := report[0]
	ascending := report[1] > report[0]
	for i, level := range report {
		if i == 0 {
			continue
		}

		if !isValidDelta(level, prev) || !isValidTrajectory(ascending, level, prev) {
			return false
		}
		prev = level
	}
	return true
}

func isValidDelta(current, prev level) bool {
	delta := absDiff(current, prev)
	return delta <= DELTA_MAX && delta >= DELTA_MIN
}

// @see {https://stackoverflow.com/a/68277627}
func absDiff(x, y level) level {
	if x < y {
		return y - x
	}
	return x - y
}

func isValidTrajectory(ascending bool, current, prev level) bool {
	if ascending {
		return current > prev
	} else {
		return current < prev
	}
}
