package day4

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	result := Solve(input)
	if result != 18 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestForwardRow(t *testing.T) {
	input := `GGGXMASGGG
ZXZMZAZSZZ
XMASGGGGGG
ZZXMZZASZZ
GGGGGGXMAS
ZZZZZZZZZZ
GXMASGXMAS
ZZZZZZZZZZ
XMASGGGGGG
ZZZZZZZZZZ`
	result := Solve(input)
	if result != 5 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestBackwardRow(t *testing.T) {
	input := `GGGSAMXGGG
ZSZAZMZXZZ
SAMXGGGGGG
ZZSAZZMXZZ
GGGGGGSAMX
ZZZZZZZZZZ
GSAMXGSAMX
ZZZZZZZZZZ
SAMXGGGGGG
ZZZZZZZZZZ`
	result := Solve(input)
	if result != 5 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestColForward(t *testing.T) {
	input := `.....
..X..
..M..
..A..
..S..`
	result := Solve(input)
	if result != 1 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestColBackward(t *testing.T) {
	input := `.....
...S.
...A.
...M.
...X.`
	result := Solve(input)
	if result != 1 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestDiagForward(t *testing.T) {
	input := `X....
.M...
..A..
...S.
.....`
	result := Solve(input)
	if result != 1 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestDiagBackward(t *testing.T) {
	input := `.....
S....
.A...
..M..
...X.`
	result := Solve(input)
	if result != 1 {
		t.Errorf("Calculated solution was not expected")
	}
}
