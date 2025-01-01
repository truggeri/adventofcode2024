package day20

import (
	"slices"
	"strings"
)

const CHEAT_LENGTH_OF_INTEREST = 100
const EMPTY_SPACE = '.'
const END = 'E'
const START = 'S'
const MOVES_TO_USE_CHEAT = 2
const THREAD_POOL = 64

type point struct {
	x, y int
}

type maze struct {
	start, end    point
	width, height uint
	emptySpaces   map[point]bool
}

func (m maze) valid(p point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < int(m.width) && p.y < int(m.height)
}

type thread struct {
	p     point
	moves []point
}

func Solve(input string) uint {
	m, sol := SolvePuzzle(input)
	cheats := Cheats(m, sol)
	var cnt uint = 0
	for k, v := range cheats {
		if k >= CHEAT_LENGTH_OF_INTEREST {
			cnt += v
		}
	}
	return cnt
}

func SolvePuzzle(input string) (maze, []point) {
	m := parseInput(input)
	sol := solve(m)
	return m, sol
}

func Cheats(m maze, sol []point) map[uint]uint {
	cheats := make(map[uint]uint)
	solution := make(map[point]uint)
	for i, p := range sol {
		solution[p] = uint(i)
	}

	for i := 1; i < int(m.width-1); i++ {
		for j := 1; j < int(m.height-1); j++ {
			p := point{x: i, y: j}
			if m.emptySpaces[p] || solution[p] > 0 {
				continue
			}

			east, south, west, north := point{x: p.x + 1, y: p.y}, point{x: p.x, y: p.y + 1}, point{x: p.x - 1, y: p.y}, point{x: p.x, y: p.y - 1}
			if (solution[east] == 0 || solution[west] == 0) && (solution[south] == 0 || solution[north] == 0) {
				continue
			}

			cheats[max(cheatLength(east, west, solution), cheatLength(south, north, solution))-MOVES_TO_USE_CHEAT] += 1
		}
	}

	return cheats
}

func parseInput(input string) maze {
	var m maze
	m.emptySpaces = make(map[point]bool)
	for j, line := range strings.Split(input, "\n") {
		m.width = uint(len(line))
		for i, r := range line {
			p := point{x: i, y: j}
			switch r {
			case EMPTY_SPACE:
				m.emptySpaces[p] = true
			case START:
				m.start = p
			case END:
				m.emptySpaces[p] = true
				m.end = p
			}
		}
		m.height = uint(j) + 1
	}
	return m
}

func solve(m maze) []point {
	seenSpaces := make(map[point]uint)
	seenSpaces[m.start] = 1
	threads := []thread{{p: m.start, moves: []point{m.start}}}
	coldStorage := make([]thread, 0)
	solution := make([]point, 0)
	for {
		nextThreads := make([]thread, 0)
		for i, t := range threads {
			if i >= THREAD_POOL {
				coldStorage = append(coldStorage, t)
				continue
			} else if t.p == m.end {
				if len(solution) == 0 || len(t.moves) < len(solution) {
					solution = t.moves
				}
			} else {
				nextThreads = append(calcNextMoves(m, t, seenSpaces), nextThreads...)
			}
		}

		if len(nextThreads) == 0 {
			if len(coldStorage) > 0 {
				x := THREAD_POOL
				if len(coldStorage) < x {
					x = len(coldStorage)
				}
				nextThreads = coldStorage[:x]
				coldStorage = slices.Delete(coldStorage, 0, x)
			} else {
				if len(solution) == 0 {
					panic("No solution with no more spaces to check")
				}
				return solution
			}
		}
		for i, t := range nextThreads {
			if len(solution) > 0 && len(t.moves) >= len(solution) {
				nextThreads = slices.Delete(nextThreads, i, i)
			}
			if seenSpaces[t.p] == 0 || uint(len(t.moves)) < seenSpaces[t.p] {
				seenSpaces[t.p] = uint(len(t.moves))
			}
		}
		threads = nextThreads
	}
}

func calcNextMoves(m maze, curr thread, seenSpaces map[point]uint) []thread {
	nextMoves := make([]thread, 0)
	possibilities := []thread{{p: point{x: curr.p.x + 1, y: curr.p.y}, moves: curr.moves},
		{p: point{x: curr.p.x, y: curr.p.y + 1}, moves: curr.moves},
		{p: point{x: curr.p.x - 1, y: curr.p.y}, moves: curr.moves},
		{p: point{x: curr.p.x, y: curr.p.y - 1}, moves: curr.moves}}
	for _, t := range possibilities {
		if m.valid(t.p) && m.emptySpaces[t.p] {
			if seenSpaces[t.p] == 0 || uint(len(t.moves))+1 < seenSpaces[t.p] {
				t.moves = append(t.moves, t.p)
				nextMoves = append(nextMoves, t)
			}
		}
	}
	return nextMoves
}

func cheatLength(a, b point, solution map[point]uint) uint {
	if solution[a] == 0 || solution[b] == 0 {
		return 0
	}

	if solution[a] > solution[b] {
		return solution[a] - solution[b]
	} else {
		return solution[b] - solution[a]
	}
}
