package day10

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	result := Solve(input)
	if result != 36 {
		t.Errorf("Calculated solution was not expected")
	}
}
