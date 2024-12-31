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

func SolvePartTwo(input string) uint {
	inputProgram := parseInputProgram(input)
	var cnt uint = 0
	for {
		cmp := parseInput(input)
		cmp.regA = cnt

		for {
			if cmp.tick() {
				break
			}
			if len(cmp.outBuffer) > 1 && strings.Join(cmp.outBuffer, ",") != inputProgram[:len(cmp.outBuffer)*2-1] {
				break
			}
		}

		if strings.Join(cmp.outBuffer, ",") == inputProgram {
			return cnt
		}
		cnt++
	}
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

func parseInputProgram(input string) string {
	lines := strings.Split(input, "\n")
	return lines[4][9:]
}
