package day7

import (
	"slices"
	"strconv"
	"strings"
)

type equation struct {
	result       uint
	coefficients []uint
}

func Solve(input string) uint {
	equations := parseInput(input)
	return solve(equations)
}

func parseInput(input string) []equation {
	result := make([]equation, 0)
	for _, line := range strings.Split(input, "\n") {
		var eqn equation
		pieces := strings.Split(line, ": ")
		if len(pieces) != 2 {
			continue
		}

		x, err := strconv.Atoi(pieces[0])
		if err != nil {
			continue
		}
		eqn.result = uint(x)

		for _, coef := range strings.Split(pieces[1], " ") {
			y, err := strconv.Atoi(coef)
			if err != nil {
				continue
			}
			eqn.coefficients = slices.Insert(eqn.coefficients, len(eqn.coefficients), uint(y))
		}

		result = slices.Insert(result, len(result), eqn)
	}
	return result
}

func solve(eqns []equation) uint {
	var result uint = 0
	for _, eqn := range eqns {
		if trueEquation(eqn) {
			result += eqn.result
		}
	}
	return result
}

func trueEquation(eqn equation) bool {
	possibilities := []uint{eqn.coefficients[0]}
	for i, coef := range eqn.coefficients {
		if i == 0 {
			continue
		}

		updatedPossibilities := []uint{}
		for _, priorPartial := range possibilities {
			sum := priorPartial + coef
			if sum <= eqn.result {
				updatedPossibilities = slices.Insert(updatedPossibilities, len(updatedPossibilities), sum)
			}
			mult := priorPartial * coef
			if mult <= eqn.result {
				updatedPossibilities = slices.Insert(updatedPossibilities, len(updatedPossibilities), mult)
			}
		}
		possibilities = updatedPossibilities
	}

	for _, poss := range possibilities {
		if poss == eqn.result {
			return true
		}
	}
	return false
}
