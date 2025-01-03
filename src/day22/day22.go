package day22

import (
	"strconv"
	"strings"
)

const ITERATIONS uint = 2000
const PRUNE_LIMIT uint = 16_777_216

func Solve(input string) uint {
	brokers := parseInput(input)
	var result uint
	for _, b := range brokers {
		result += secretNumber(b, ITERATIONS)
	}
	return result
}

func parseInput(input string) []uint {
	result := make([]uint, 0)
	for _, line := range strings.Split(input, "\n") {
		x, _ := strconv.Atoi(line)
		result = append(result, uint(x))
	}
	return result
}

func secretNumber(seed, iterations uint) uint {
	result := seed
	for range iterations {
		result = iterate(result)
	}
	return result
}

func iterate(secret uint) uint {
	step1 := prune(mix(secret, secret<<6))
	step2 := prune(mix(step1, step1>>5))
	step3 := prune(mix(step2, step2<<11))
	return step3
}

func mix(secret, v uint) uint {
	return secret ^ v
}

func prune(v uint) uint {
	return v & (PRUNE_LIMIT - 1)
}
