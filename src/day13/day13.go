package day13

import (
	"strconv"
	"strings"
)

const BUTTON_PUSH_LIMIT = 100
const BUTTON_A_COST = 3
const BUTTON_B_COST = 1

type point struct {
	x, y int
}

type machine struct {
	buttonA, buttonB point
	prize            point
}

func Solve(input string) (uint, uint) {
	mm := parse(input)
	var prizes, cost uint
	for _, m := range mm {
		win, c := solveMachine(m)
		if win {
			prizes++
			cost += c
		}
	}
	return prizes, cost
}

func parse(input string) []machine {
	machines := make([]machine, 0)
	currentMachine := machine{}

	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "Button ") {
			data := strings.Split(line[12:], ", Y+")
			if len(data) != 2 {
				continue
			}

			x, _ := strconv.Atoi(data[0])
			y, _ := strconv.Atoi(data[1])
			point := point{x, y}
			buttonA := line[7] == 'A'

			if buttonA {
				currentMachine.buttonA = point
			} else {
				currentMachine.buttonB = point
			}
		} else if strings.Contains(line, "Prize: ") {
			data := strings.Split(line[9:], ", Y=")
			if len(data) != 2 {
				continue
			}
			x, _ := strconv.Atoi(data[0])
			currentMachine.prize.x = x
			y, _ := strconv.Atoi(data[1])
			currentMachine.prize.y = y
		} else {
			machines = append(machines, currentMachine)
		}
	}
	machines = append(machines, currentMachine)
	return machines
}

func solveMachine(m machine) (bool, uint) {
	a, b := cramersRule([2][2]int{{m.buttonA.x, m.buttonB.x}, {m.buttonA.y, m.buttonB.y}}, [2][1]int{{m.prize.x}, {m.prize.y}})
	if a == 0 && b == 0 {
		return manualSolution(m)
	}
	return validSolution(m, a, b), uint(a*BUTTON_A_COST + b*BUTTON_B_COST)
}

// @see{https://en.wikipedia.org/wiki/Cramer's_rule#Explicit_formulas_for_small_systems}
func cramersRule(coef [2][2]int, column [2][1]int) (int, int) {
	coefMatrix := coef
	det := determinate(coefMatrix)
	if det == 0 {
		return 0, 0
	}

	x := determinate([2][2]int{{column[0][0], coef[0][1]}, {column[1][0], coef[1][1]}}) / det
	y := determinate([2][2]int{{coef[0][0], column[0][0]}, {coef[1][0], column[1][0]}}) / det
	return x, y
}

// @see{https://en.wikipedia.org/wiki/Determinant}
func determinate(input [2][2]int) int {
	return input[0][0]*input[1][1] - input[0][1]*input[1][0]
}

func validSolution(m machine, a, b int) bool {
	positiveOnly := (a > 0 && b >= 0) || (a >= 0 && b > 0)
	withinLimits := a <= 100 && b <= 100
	validAnswer := (a*m.buttonA.x+b*m.buttonB.x == m.prize.x) && (a*m.buttonA.y+b*m.buttonB.y == m.prize.y)

	return positiveOnly && withinLimits && validAnswer
}

func manualSolution(m machine) (bool, uint) {
	a, costa := linearSolution(m.buttonA, m.prize, BUTTON_A_COST)
	b, costb := linearSolution(m.buttonB, m.prize, BUTTON_B_COST)

	if a && b {
		return true, min(costa, costb)
	} else if a {
		return a, costa
	} else if b {
		return b, costb
	}
	return false, 0
}

func linearSolution(a, p point, cost int) (bool, uint) {
	zeroXPresses := p.x == 0 && a.x == 0
	zeroYPresses := p.y == 0 && a.y == 0
	x := zeroXPresses || p.x%a.x == 0
	y := zeroYPresses || p.y%a.y == 0
	solutionExists := x && y && (zeroXPresses || zeroYPresses || p.x/a.x == p.y/a.y)
	if !solutionExists {
		return false, 0
	}

	var presses int
	if zeroXPresses {
		presses = p.y / a.y
	} else {
		presses = p.x / a.x
	}
	return true, uint(presses * cost)
}
