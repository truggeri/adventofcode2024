package day17

import (
	"strconv"
	"strings"
)

func Solve(input string) string {
	cmp := parseInput(input)
	for {
		if cmp.tick() {
			break
		}
	}
	return strings.Join(cmp.outBuffer, ",")
}

func parseInput(input string) computer {
	lines := strings.Split(input, "\n")
	a, _ := strconv.Atoi(lines[0][12:])
	b, _ := strconv.Atoi(lines[1][12:])
	c, _ := strconv.Atoi(lines[2][12:])
	cmp := computer{regA: uint(a), regB: uint(b), regC: uint(c)}
	for _, instr := range strings.Split(lines[4][9:], ",") {
		i, _ := strconv.Atoi(instr)
		cmp.instructions = append(cmp.instructions, instruction(i))
	}
	return cmp
}
