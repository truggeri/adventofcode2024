package day1

import (
	"slices"
	"strconv"
	"strings"
)

// This is a bit fragile, but the spec doesn't indicate it will ever change
const INPUT_SEPARATOR = "   "

func Solve(input string) uint {
	return solve(parseInput(input))
}

func PartTwo(input string) uint {
	left, right := parseInput(input)
	occurances := make(map[uint]uint)
	for _, r := range right {
		occurances[r] += 1
	}
	var result uint = 0
	for _, l := range left {
		result += l * occurances[l]
	}
	return result
}

func parseInput(input string) ([]uint, []uint) {
	var a, b []uint
	for _, line := range strings.Split(input, "\n") {
		values := strings.Split(line, INPUT_SEPARATOR)
		if len(values) != 2 {
			continue
		}

		x, _ := strconv.Atoi(values[0])
		a = append(a, uint(x))
		y, _ := strconv.Atoi(values[1])
		b = append(b, uint(y))
	}
	return a, b
}

func solve(a, b []uint) uint {
	slices.Sort(a)
	slices.Sort(b)
	var sum uint = 0
	for i := range a {
		sum += absDiffUint(a[i], b[i])
	}
	return sum
}

// @see {https://stackoverflow.com/a/68277627}
func absDiffUint(x, y uint) uint {
	if x < y {
		return y - x
	}
	return x - y
}
