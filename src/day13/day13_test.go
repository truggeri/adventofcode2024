package day13

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`
	prizes, tokens := Solve(input)
	if prizes != 2 || tokens != 480 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestMultipleA(t *testing.T) {
	input := `Button A: X+50, Y+25
Button B: X+10, Y+5
Prize: X=100, Y=50`
	prizes, tokens := Solve(input)
	if prizes != 1 || tokens != 6 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestMultipleB(t *testing.T) {
	input := `Button A: X+50, Y+25
Button B: X+20, Y+10
Prize: X=100, Y=50`
	prizes, tokens := Solve(input)
	if prizes != 1 || tokens != 5 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestNoSolutionDetZero(t *testing.T) {
	input := `Button A: X+4, Y+0
Button B: X+3, Y+0
Prize: X=5, Y=0`
	prizes, tokens := Solve(input)
	if prizes != 0 || tokens != 0 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestNoSolutionDetPositive(t *testing.T) {
	input := `Button A: X+3, Y+1
Button B: X+2, Y+1
Prize: X=5, Y=1`
	prizes, tokens := Solve(input)
	if prizes != 0 || tokens != 0 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestNoSolutionDetNeg(t *testing.T) {
	input := `Button A: X+2, Y+1
Button B: X+3, Y+1
Prize: X=5, Y=1`
	prizes, tokens := Solve(input)
	if prizes != 0 || tokens != 0 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestSolutionOnXAxis(t *testing.T) {
	input := `Button A: X+2, Y+0
Button B: X+3, Y+0
Prize: X=4, Y=0`
	prizes, tokens := Solve(input)
	if prizes != 1 || tokens != 6 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestSolutionOnYAxis(t *testing.T) {
	input := `Button A: X+0, Y+4
Button B: X+0, Y+3
Prize: X=0, Y=9`
	prizes, tokens := Solve(input)
	if prizes != 1 || tokens != 3 {
		t.Errorf("Calculated solution was not expected")
	}
}
