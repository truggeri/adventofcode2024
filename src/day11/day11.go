package day11

import (
	"slices"
	"strings"
)

const ZERO_RUNE = '0'
const BLINKS = 25

var MULTIPLE = []uint{4, 2, 0, 2}

type stone []uint

func Solve(input string) uint {
	initial := parseInput(input)
	return solve(initial, BLINKS)
}

func parseInput(input string) []stone {
	result := make([]stone, 0)
	for i, num := range strings.Split(input, " ") {
		result = slices.Insert(result, i, make(stone, 0))
		for _, chr := range num {
			result[i] = append(result[i], runeToUInt(chr))
		}
		slices.Reverse(result[i])
	}
	return result
}

func runeToUInt(r rune) uint {
	return uint(r - ZERO_RUNE)
}

func solve(input []stone, blinks uint) uint {
	curr := input
	for range blinks {
		curr = blink(curr)
	}
	return uint(len(curr))
}

func blink(input []stone) []stone {
	result := make([]stone, 0)
	for _, stn := range input {
		if isZero(stn) {
			result = append(result, stone{1})
		} else if isEvenDigits(stn) {
			splt := split(stn)
			result = append(result, splt[0], splt[1])
		} else {
			result = append(result, multiply(stn))
		}
	}
	return result
}

func isZero(s stone) bool {
	return len(s) == 1 && s[0] == 0
}

func isEvenDigits(s stone) bool {
	return len(s)%2 == 0
}

func split(input stone) [2]stone {
	digits := len(input)
	var result [2]stone
	result[0] = zeroTrim(input[digits/2 : digits])
	result[1] = zeroTrim(input[0 : digits/2])
	return result
}

func zeroTrim(input stone) stone {
	lastDigit := 0
	for i, d := range input {
		if d != 0 {
			lastDigit = i
		}
	}
	return input[0 : lastDigit+1]
}

// @see{https://en.wikipedia.org/wiki/Multiplication_algorithm#Example}
func multiply(input stone) stone {
	result := make(stone, len(input)+len(MULTIPLE))
	var carry uint
	for i := range input {
		carry = 0
		for a := range MULTIPLE {
			result[a+i] += carry + MULTIPLE[a]*input[i]
			carry = result[a+i] / 10
			result[a+i] = result[a+i] % 10
		}
		result[i+len(MULTIPLE)] = carry
	}
	return zeroTrim(result)
}
