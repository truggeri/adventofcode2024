package day3

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const MUL_PREFIX = "mul("
const MUL_LEN = len(MUL_PREFIX)
const MUL_RE = `mul\(([0-9]{1,3}),([0-9]{1,3})\)`

type mul struct {
	x, y uint
}

func (m mul) Eval() uint {
	return m.x * m.y
}

func Solve(input string) uint {
	return accumulate(parseInput(input))
}

func parseInput(input string) []mul {
	var results []mul
	re := regexp.MustCompile(MUL_RE)
	matches := re.FindAllString(input, -1)
	if matches == nil {
		return results
	}

	for _, match := range matches {
		newValue := parseMatch(match)
		results = slices.Insert(results, len(results), newValue)
	}

	return results
}

func parseMatch(input string) mul {
	numbers := strings.Split(input[MUL_LEN:len(input)-1], ",")
	if len(numbers) != 2 {
		return mul{}
	}

	x, err := strconv.Atoi(numbers[0])
	if err != nil {
		return mul{}
	}
	y, err := strconv.Atoi(numbers[1])
	if err != nil {
		return mul{}
	}

	return mul{x: uint(x), y: uint(y)}
}

func accumulate(input []mul) uint {
	var result uint = 0
	for _, mul := range input {
		result += mul.Eval()
	}
	return result
}
