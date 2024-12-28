package day9

import (
	"math"
)

const ZERO_RUNE = '0'
const EMPTY_RUNE = '.'

type file struct {
	empty bool
	digit uint
}

func Solve(input string) uint64 {
	return checksum(compress(expand(input)))
}

func expand(input string) []file {
	result := make([]file, 0)
	blankSpace := false
	var id uint = 0
	for _, fileSize := range input {
		for i := 0; i < runeToInt(fileSize); i++ {
			result = append(result, file{empty: blankSpace, digit: id})
		}
		if !blankSpace {
			id += 1
		}
		blankSpace = !blankSpace
	}
	return result
}

func runeToInt(r rune) int {
	return int(r - ZERO_RUNE)
}

func compress(input []file) []file {
	result := input
	ptr, tailPtr := 0, len(input)-1
	for {
		if !result[tailPtr].empty || tailPtr <= 0 {
			break
		}
		tailPtr--
	}

	for {
		if ptr >= tailPtr {
			break
		}
		if !result[ptr].empty {
			ptr++
			continue
		}
		result[ptr] = result[tailPtr]
		result[tailPtr] = file{empty: true}
		ptr++
		for {
			tailPtr--
			if !result[tailPtr].empty {
				break
			}
		}
	}
	return result
}

func checksum(input []file) uint64 {
	result := uint64(0)
	for i, d := range input {
		if d.empty {
			break
		}
		val := uint64(i * int(d.digit))
		if math.MaxUint64-val < result {
			panic("overflow")
		}
		result += val
	}
	return result
}
