package day9

import (
	"strconv"
)

const ZERO_RUNE = '0'
const EMPTY_RUNE = '.'
const EMPTY_STRING = string(EMPTY_RUNE)

func Solve(input string) int {
	return checksum(compress(expand(input)))
}

func expand(input string) string {
	var result string
	blankSpace := false
	id := 0
	for _, fileSize := range input {
		for i := 0; i < runeToInt(fileSize); i++ {
			if blankSpace {
				result += EMPTY_STRING
			} else {
				result += strconv.Itoa(id)
			}
		}
		if !blankSpace {
			id += 1
		}
		blankSpace = !blankSpace
	}
	return result
}

func compress(input string) string {
	result := []byte(input)
	ptr, tailPtr := 0, len(input)-1
	for {
		if result[tailPtr] != EMPTY_RUNE || tailPtr <= 0 {
			break
		}
		tailPtr--
	}

	for {
		if ptr >= tailPtr {
			break
		}
		if result[ptr] != EMPTY_RUNE {
			ptr++
			continue
		}
		result[ptr] = result[tailPtr]
		result[tailPtr] = EMPTY_RUNE
		ptr++
		for {
			tailPtr--
			if result[tailPtr] != EMPTY_RUNE {
				break
			}
		}
	}
	return string(result)
}

func checksum(input string) int {
	result := 0
	for i, d := range input {
		if d == EMPTY_RUNE {
			break
		}
		result += i * runeToInt(d)
	}
	return result
}

func runeToInt(r rune) int {
	return int(r - ZERO_RUNE)
}
