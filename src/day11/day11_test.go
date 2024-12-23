package day11

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := `125 17`
	result := Solve(input)
	if result != 55312 {
		t.Errorf("Calculated solution was not expected")
	}
}
