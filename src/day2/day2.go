package day2

import (
	"slices"
	"strconv"
	"strings"
)

const DELTA_MAX = 3
const DELTA_MIN = 1

type level uint
type report []level

func Solve(input string) uint {
	return solve(parseInput(input), 0)
}

func PartTwo(input string) uint {
	return solve(parseInput(input), 1)
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

func solve(input []report, tolerance uint) uint {
	var result uint = 0
	for _, report := range input {
		if isValidReport(report, tolerance) {
			result += 1
		}
	}
	return result
}

func isValidReport(report report, tolerance uint) bool {
	if len(report) < 2 {
		return true
	}

	var violations uint = 0
	ascending := report[1] > report[0]
	i := 0
	for {
		i++
		if i > len(report)-1 {
			break
		}

		if !isValidDelta(report[i], report[i-1]) || !isValidTrajectory(ascending, report[i], report[i-1]) {
			violations++
			if violations > tolerance {
				return false
			}

			selfRemoved := slices.Concat(report[0:i], report[i+1:])
			valid := isValidReport(selfRemoved, tolerance-violations)
			if valid {
				return true
			}

			prevRemoved := slices.Concat(report[0:i-1], report[i:])
			valid = isValidReport(prevRemoved, tolerance-violations)
			if valid {
				return true
			}

			if i == 1 {
				return false
			}
			twoBackRemoved := slices.Concat(report[0:i-2], report[i-1:])
			valid = isValidReport(twoBackRemoved, tolerance-violations)
			return valid
		}
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
