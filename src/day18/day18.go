package day18

import (
	"slices"
	"strconv"
	"strings"
)

const MAX_THREADS_PER_LOOP = 32

type point struct {
	x, y int
}

type maze struct {
	walls  map[point]bool
	start  point
	end    point
	width  int
	height int
}

func (m maze) valid(p point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < m.width && p.y < m.height
}

type thread struct {
	p     point
	moves uint
}

func Solve(input string, size int, square int) uint {
	m := maze{width: square, height: square, start: point{0, 0}, end: point{x: square - 1, y: square - 1}, walls: make(map[point]bool)}
	blocks := parseInput(input)
	for i := 0; i < size; i++ {
		m.walls[blocks[i]] = true
	}
	return solve(m)
}

func parseInput(input string) []point {
	result := make([]point, 0)
	for _, line := range strings.Split(input, "\n") {
		d := strings.Split(line, ",")
		if len(d) != 2 {
			continue
		}

		x, _ := strconv.Atoi(d[0])
		y, _ := strconv.Atoi(d[1])
		result = append(result, point{x: x, y: y})
	}
	return result
}

func solve(m maze) uint {
	seenSpaces := make(map[point]uint)
	seenSpaces[m.start] = 1
	threads := []thread{{p: m.start}}
	coldStorage := make([]thread, 0)
	var solution uint = 0
	loops := 0
	var threadCount uint
	for {
		nextThreads := make([]thread, 0)
		threadCount += uint(len(threads))
		for i, t := range threads {
			if i >= MAX_THREADS_PER_LOOP {
				coldStorage = append(coldStorage, t)
				continue
			} else if t.p == m.end {
				if solution == 0 || t.moves < solution {
					solution = t.moves
				}
			} else {
				nextThreads = append(calcNextMoves(m, t, seenSpaces), nextThreads...)
			}
		}

		if len(nextThreads) == 0 {
			if len(coldStorage) > 0 {
				x := MAX_THREADS_PER_LOOP
				if len(coldStorage) < x {
					x = len(coldStorage)
				}
				nextThreads = coldStorage[:x]
				coldStorage = slices.Delete(coldStorage, 0, x)
			} else {
				if solution == 0 {
					panic("No solution with no more spaces to check")
				}
				return solution
			}
		}
		for i, t := range nextThreads {
			if solution > 0 && t.moves >= solution {
				nextThreads = slices.Delete(nextThreads, i, i)
			}
			if seenSpaces[t.p] == 0 || t.moves < seenSpaces[t.p] {
				seenSpaces[t.p] = t.moves
			}
		}
		threads = nextThreads
		loops++
	}
}

func calcNextMoves(m maze, curr thread, seenSpaces map[point]uint) []thread {
	nextMoves := make([]thread, 0)
	possibilities := []thread{{p: point{x: curr.p.x + 1, y: curr.p.y}, moves: curr.moves},
		{p: point{x: curr.p.x, y: curr.p.y + 1}, moves: curr.moves},
		{p: point{x: curr.p.x - 1, y: curr.p.y}, moves: curr.moves},
		{p: point{x: curr.p.x, y: curr.p.y - 1}, moves: curr.moves}}
	for _, t := range possibilities {
		if m.valid(t.p) && !m.walls[t.p] {
			t.moves++
			if seenSpaces[t.p] == 0 || t.moves < seenSpaces[t.p] {
				nextMoves = append(nextMoves, t)
			}
		}
	}
	return nextMoves
}
