package day15

const WALL = '#'
const BOX = 'O'
const ROBOT = '@'
const EMPTY = '.'

const UP = '^'
const DOWN = 'v'
const LEFT = '<'
const RIGHT = '>'

const INVALID_RUNE = '-'

type point struct {
	x, y int
}

type puzzle struct {
	width, height int
	warehouse     map[point]bool
	boxes         map[point]bool
	instructions  []rune
	instrPtr      int
	robot         point
}

type moveFunc func(point) point

func puzzleBuilder() puzzle {
	var puz puzzle
	puz.warehouse = make(map[point]bool)
	puz.boxes = make(map[point]bool)
	puz.instructions = make([]rune, 0)
	puz.instrPtr = 0
	return puz
}

func (puz *puzzle) get(p point) rune {
	if !puz.validPoint(p) {
		return INVALID_RUNE
	} else if puz.warehouse[p] {
		return WALL
	} else if puz.boxes[p] {
		return BOX
	} else if puz.robot == p {
		return ROBOT
	}
	return EMPTY
}

func (puz *puzzle) set(p point, r rune) {
	current := puz.get(p)
	if current == r || current == INVALID_RUNE {
		return
	}

	switch current {
	case WALL:
		delete(puz.warehouse, p)
		break
	case BOX:
		delete(puz.boxes, p)
		break
	case ROBOT:
		puz.robot = point{-1, -1}
		break
	}

	switch r {
	case WALL:
		puz.warehouse[p] = true
		break
	case BOX:
		puz.boxes[p] = true
		break
	case ROBOT:
		puz.robot = p
		break
	}
}

func (puz *puzzle) validPoint(p point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < puz.width && p.y < puz.height
}

func (puz *puzzle) nextInstruction() (rune, bool) {
	if puz.instrPtr >= len(puz.instructions) {
		return INVALID_RUNE, false
	}

	puz.instrPtr++
	return puz.instructions[puz.instrPtr-1], true
}

func (puz *puzzle) mover(direction rune) moveFunc {
	var mover moveFunc
	switch direction {
	case UP:
		mover = func(p point) point {
			return point{p.x, p.y - 1}
		}
	case DOWN:
		mover = func(p point) point {
			return point{p.x, p.y + 1}
		}
	case LEFT:
		mover = func(p point) point {
			return point{p.x - 1, p.y}
		}
	case RIGHT:
		mover = func(p point) point {
			return point{p.x + 1, p.y}
		}
	}
	return mover
}

func (puz *puzzle) move(mover moveFunc) int {
	rbt := puz.robot
	moves := 0
	for {
		rbt = mover(rbt)
		moves++
		if !puz.validPoint(rbt) {
			return 0
		} else if puz.warehouse[rbt] {
			return 0
		} else if !puz.boxes[rbt] {
			return moves
		}
	}
}

func (puz *puzzle) gpsSum() uint {
	var sum uint = 0
	for p := range puz.boxes {
		sum += uint(p.x + 100*p.y)
	}
	return sum
}
