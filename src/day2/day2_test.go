package day2

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	result := Solve(input)
	if result != 2 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestSamplePartTwo(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	result := PartTwo(input)
	if result != 4 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestPartTwoFirstTooLarge(t *testing.T) {
	input := `1 5 6 7 8`
	result := PartTwo(input)
	if result != 1 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestPartTwoFirstThrowsTraj(t *testing.T) {
	input := `4 5 4 3 2`
	result := PartTwo(input)
	if result != 1 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestPartTwoLastTooLarge(t *testing.T) {
	input := `1 2 3 4 9`
	result := PartTwo(input)
	if result != 1 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestPartTwoLastThrowsTraj(t *testing.T) {
	input := `5 4 3 2 3`
	result := PartTwo(input)
	if result != 1 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestPartTwo2ndLastTooLarge(t *testing.T) {
	input := `9 8 7 4 5`
	result := PartTwo(input)
	if result != 1 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestPartTwo2ndLastThrowsTraj(t *testing.T) {
	input := `5 4 3 4 1`
	result := PartTwo(input)
	if result != 1 {
		t.Errorf("Calculated solution was not expected")
	}
}
