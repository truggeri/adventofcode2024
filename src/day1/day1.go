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

func parseInput(input string) ([]uint, []uint) {
	var a, b []uint
	for _, line := range strings.Split(input, "\n") {
		values := strings.Split(line, INPUT_SEPARATOR)
		if len(values) != 2 {
			continue
		}

		x, _ := strconv.Atoi(values[0])
		a = slices.Concat(a, []uint{uint(x)})
		y, _ := strconv.Atoi(values[1])
		b = slices.Concat(b, []uint{uint(y)})
	}
	return a, b
}

func solve(a, b []uint) uint {
	slices.Sort(a)
	slices.Sort(b)
	var sum uint = 0
	for i := range a {
		sum = sum + absDiffUint(a[i], b[i])
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
