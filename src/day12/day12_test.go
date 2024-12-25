package day12

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	result := Solve(input)
	if result != 140 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestSample2(t *testing.T) {
	input := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	result := Solve(input)
	if result != 772 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestSample3(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	result := Solve(input)
	if result != 1930 {
		t.Errorf("Calculated solution was not expected")
	}
}
