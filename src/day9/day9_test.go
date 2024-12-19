package day9

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := `2333133121414131402`
	result := Solve(input)
	if result != 1928 {
		t.Errorf("Calculated solution was not expected")
	}
}
