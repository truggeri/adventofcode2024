package day7

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
	result := Solve(input)
	if result != 3749 {
		t.Errorf("Calculated solution was not expected")
	}
}
