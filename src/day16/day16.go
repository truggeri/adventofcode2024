package day16

import (
	"slices"
	"strings"
)

const EMPTY = '.'
const END = 'E'
const START = 'S'

type direction uint8

const (
	EAST direction = iota
	SOUTH
	WEST
	NORTH
)

type point struct {
	x, y int
}

func (p point) valid(w, h int) bool {
	return p.x >= 0 && p.y >= 0 && p.x < w && p.y < h
}

type maze struct {
	emptySpaces map[point]bool
	start       point
	end         point
	width       int
	height      int
}

type thread struct {
	p     point
	dir   direction
	moves uint
	turns uint
}

func (t thread) cost() uint {
	return t.moves + 1000*t.turns
}

func Solve(input string) uint {
	m := parseInput(input)
	return solve(m)
}

func parseInput(input string) maze {
	var result maze
	result.emptySpaces = make(map[point]bool)

	result.height = len(input)
	for j, line := range strings.Split(input, "\n") {
		result.width = len(line)
		for i, r := range line {
			if r == EMPTY || r == START || r == END {
				result.emptySpaces[point{x: i, y: j}] = true
				if r == START {
					result.start = point{x: i, y: j}
				} else if r == END {
					result.end = point{x: i, y: j}
				}
			}
		}
	}
	return result
}

func solve(m maze) uint {
	seenSpaces := make(map[point]uint)
	seenSpaces[m.start] = 1
	threads := []thread{{p: m.start, dir: EAST}}
	var solution uint = 0

	for {
		nextThreads := make([]thread, 0)
		for _, t := range threads {
			if t.p == m.end {
				if solution == 0 || t.cost() < solution {
					solution = t.cost()
				}
			} else {
				nextThreads = append(nextThreads, calcNextMoves(m, t, seenSpaces)...)
			}
		}

		if len(nextThreads) == 0 {
			if solution == 0 {
				panic("No solution with no more spaces to check")
			}
			return solution
		}
		for i, t := range nextThreads {
			if solution > 0 && t.cost() >= solution {
				nextThreads = slices.Delete(nextThreads, i, i)
			}
			if seenSpaces[t.p] == 0 || t.cost() < seenSpaces[t.p] {
				seenSpaces[t.p] = t.cost()
			}
		}
		threads = nextThreads
	}
}

func calcNextMoves(m maze, curr thread, seenSpaces map[point]uint) []thread {
	nextMoves := make([]thread, 0)
	possibilities := []thread{{p: point{x: curr.p.x + 1, y: curr.p.y}, dir: EAST, moves: curr.moves, turns: curr.turns},
		{p: point{x: curr.p.x, y: curr.p.y + 1}, dir: SOUTH, moves: curr.moves, turns: curr.turns},
		{p: point{x: curr.p.x - 1, y: curr.p.y}, dir: WEST, moves: curr.moves, turns: curr.turns},
		{p: point{x: curr.p.x, y: curr.p.y - 1}, dir: NORTH, moves: curr.moves, turns: curr.turns}}
	for i, t := range possibilities {
		if t.p.valid(m.width, m.height) && m.emptySpaces[t.p] {
			t.moves++
			if i%2 != int(curr.dir)%2 {
				t.turns++
			} else if int(curr.dir) != i {
				t.turns += 2
			}
			if seenSpaces[t.p] == 0 || t.cost() < seenSpaces[t.p] {
				nextMoves = append(nextMoves, t)
			}
		}
	}
	return nextMoves
}
