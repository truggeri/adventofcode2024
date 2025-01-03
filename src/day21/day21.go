package day21

import (
	"strconv"
	"strings"
)

const DIR_PAD_ROBOTS = 2

type numPad uint8

const (
	NUM_0 numPad = 1
	NUM_A numPad = 2
	NUM_1 numPad = 3
	NUM_2 numPad = 4
	NUM_3 numPad = 5
	NUM_4 numPad = 6
	NUM_5 numPad = 7
	NUM_6 numPad = 8
	NUM_7 numPad = 9
	NUM_8 numPad = 10
	NUM_9 numPad = 11
)

type dirPad uint8

const (
	DIR_LEFT  dirPad = 0
	DIR_DOWN  dirPad = 1
	DIR_RIGHT dirPad = 2
	DIR_UP    dirPad = 4
	DIR_A     dirPad = 5
)

func Solve(input string) uint {
	var result uint
	codes := strings.Split(input, "\n")
	for _, code := range codes {
		seq := directionSequence(code, DIR_PAD_ROBOTS)
		result += uint(len(seq)) * codeToNum(code)
	}
	return result
}

func directionSequence(code string, n uint) []dirPad {
	r := solveDirPad(code)
	for range n {
		r = sovleNumPad(r)
	}
	return r
}

func solveDirPad(code string) []dirPad {
	result := make([]dirPad, 0)
	curr := NUM_A
	for _, r := range code {
		n := runeToNum(r)
		result = append(result, numPadMoves(curr, n)...)
		curr = n
	}
	return result
}

func runeToNum(r rune) numPad {
	if r == 'A' {
		return NUM_A
	} else if r == '0' {
		return NUM_0
	} else if (r-'0' >= 1) && (r-'0' <= 9) {
		return NUM_1 + numPad(r-'1')
	}
	panic("Found rune that was not on num pad")
}

func numPadMoves(starting, ending numPad) []dirPad {
	f := func(s, e, x, y int) bool {
		starting := numPad(s)
		ending := numPad(e)
		xThroughBlank := (starting == NUM_0 || starting == NUM_A) && (ending == NUM_1 || ending == NUM_4 || ending == NUM_7)
		return !xThroughBlank && x < 0
	}
	return moves(int(starting), int(ending), f)
}

func moves(starting, ending int, moveXFirst func(int, int, int, int) bool) []dirPad {
	result := make([]dirPad, 0)
	x, y := ending%3-starting%3, starting/3-ending/3
	leftRight, upDown := DIR_RIGHT, DIR_DOWN
	if x < 0 {
		leftRight = DIR_LEFT
	}
	if y < 0 {
		upDown = DIR_UP
	}
	xFirst := moveXFirst(starting, ending, x, y)
	if xFirst {
		for range abs(x) {
			result = append(result, leftRight)
		}
	}
	for range abs(y) {
		result = append(result, upDown)
	}
	if !xFirst {
		for range abs(x) {
			result = append(result, leftRight)
		}
	}
	result = append(result, DIR_A)
	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func sovleNumPad(code []dirPad) []dirPad {
	result := make([]dirPad, 0)
	curr := DIR_A
	for _, d := range code {
		result = append(result, dirPadMoves(curr, d)...)
		curr = d
	}
	return result
}

func dirPadMoves(starting dirPad, ending dirPad) []dirPad {
	f := func(s, e, x, y int) bool {
		return y < 0
	}
	return moves(int(starting), int(ending), f)
}

func codeToNum(code string) uint {
	x, _ := strconv.Atoi(code[:len(code)-1])
	return uint(x)
}
