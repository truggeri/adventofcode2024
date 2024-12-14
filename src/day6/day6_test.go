package day6

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	result := Solve(input)
	if result != 41 {
		t.Errorf("Calculated solution was not expected")
	}
}
