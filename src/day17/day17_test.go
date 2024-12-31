package day17

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	if Solve(input) != "4,6,3,5,6,3,5,2,1,0" {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestPart1(t *testing.T) {
	input := `Register A: 41644071
Register B: 0
Register C: 0

Program: 2,4,1,2,7,5,1,7,4,4,0,3,5,5,3,0`
	if Solve(input) != "3,1,5,3,7,4,2,7,5" {
		t.Errorf("Calculated solution was not expected")
	}
}
