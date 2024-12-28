package day1

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	result := Solve(input)
	if result != 11 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestSamplePartTwo(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	result := PartTwo(input)
	if result != 31 {
		t.Errorf("Calculated solution was not expected")
	}
}
