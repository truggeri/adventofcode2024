package day15

import (
	"strings"
)

func Solve(input string) uint {
	puz := parseInput(input)
	for {
		if !iterate(&puz) {
			break
		}
	}
	return puz.gpsSum()
}

func parseInput(input string) puzzle {
	puz := puzzleBuilder()
	instr := false
	for j, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			instr = true
			continue
		}

		if !instr {
			if puz.width < len(line) {
				puz.width = len(line)
			}
			if puz.height <= j {
				puz.height = j + 1
			}

			for i, r := range line {
				p := point{x: i, y: j}
				if r == WALL {
					puz.warehouse[p] = true
				} else if r == BOX {
					puz.boxes[p] = true
				} else if r == ROBOT {
					puz.robot = p
				}
			}
		} else {
			for _, r := range line {
				puz.instructions = append(puz.instructions, r)
			}
		}
	}
	return puz
}

func iterate(puz *puzzle) bool {
	nextInstr, ok := puz.nextInstruction()
	if !ok {
		return false
	}

	mover := puz.mover(nextInstr)
	m := puz.move(mover)
	if m == 0 {
		return true
	}

	prev := EMPTY
	p := puz.robot
	for range m + 1 {
		curr := puz.get(p)
		puz.set(p, prev)
		prev = curr
		p = mover(p)
	}
	return true
}
