package day21

import (
	"testing"
)

func runeToDir(r rune) dirPad {
	switch r {
	case '<':
		return DIR_LEFT
	case 'v':
		return DIR_DOWN
	case '>':
		return DIR_RIGHT
	case '^':
		return DIR_UP
	case 'A':
		return DIR_A
	}
	return 0
}

func checkSeq(t *testing.T, expected string, result []dirPad) {
	for i, r := range expected {
		if runeToDir(r) != result[i] {
			t.Errorf("Value %d at index %d did not match expected", result[i], i)
		}
	}
}

func TestSample(t *testing.T) {
	input := `029A
980A
179A
456A
379A`
	result := Solve(input)
	if result != 126384 {
		t.Errorf("Calculated solution (%d) was not expected", result)
	}
}

func TestComplexity029A(t *testing.T) {
	input := `029A`
	result := Solve(input)
	if result != 1972 {
		t.Errorf("Calculated solution (%d) was not expected", result)
	}
}

func TestSeq029A(t *testing.T) {
	input := `029A`
	result := directionSequence(input, 2)
	if len(result) > 68 {
		t.Errorf("Calculated solution (%d) was not expected", len(result))
	}
	checkSeq(t, "v<A<AA>>^AvAA<^A>Av<<A>>^AvA^Av<<A>>^AAv<A>A^A<A>Av<A<A>>^AAAvA<^A>A", result)
}

func TestComplexity980A(t *testing.T) {
	input := `980A`
	result := Solve(input)
	if result != 58800 {
		t.Errorf("Calculated solution (%d) was not expected", result)
	}
}

func TestSeq980A(t *testing.T) {
	input := `980A`
	result := directionSequence(input, 2)
	if len(result) > 60 {
		t.Errorf("Calculated solution (%d) was not expected", len(result))
	}
	checkSeq(t, "v<<A>>^AAAvA^Av<A<AA>>^AvAA<^A>Av<A<A>>^AAAvA<^A>Av<A>^A<A>A", result)
}

func TestComplexity179A(t *testing.T) {
	input := `179A`
	result := Solve(input)
	if result != 12172 {
		t.Errorf("Calculated solution (%d) was not expected", result)
	}
}

func TestSeq179A(t *testing.T) {
	input := `179A`
	result := directionSequence(input, 2)
	if len(result) > 68 {
		t.Errorf("Calculated solution (%d) was not expected", len(result))
	}
	checkSeq(t, "v<<A>>^Av<A<A>>^AAvAA<^A>Av<<A>>^AAvA^Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A", result)
}

func TestComplexity456A(t *testing.T) {
	input := `456A`
	result := Solve(input)
	if result != 29184 {
		t.Errorf("Calculated solution (%d) was not expected", result)
	}
}

func TestSeq456A(t *testing.T) {
	input := `456A`
	result := directionSequence(input, 2)
	if len(result) > 64 {
		t.Errorf("Calculated solution (%d) was not expected", len(result))
	}
	checkSeq(t, "v<<A>>^AAv<A<A>>^AAvAA<^A>Av<A>^A<A>Av<A>^A<A>Av<A<A>>^AAvA<^A>A", result)
}

func TestComplexity379A(t *testing.T) {
	input := `379A`
	result := Solve(input)
	if result != 24256 {
		t.Errorf("Calculated solution (%d) was not expected", result)
	}
}

func TestSeq379A(t *testing.T) {
	input := `379A`
	result := directionSequence(input, 2)
	if len(result) > 64 {
		t.Errorf("Calculated solution (%d) was not expected", len(result))
	}
	checkSeq(t, "v<<A>>^AvA^Av<A<AA>>^AAvA<^A>AAvA^Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A", result)
}

func TestPart1(t *testing.T) {
	input := `869A
180A
596A
965A
973A`
	result := Solve(input)
	if result != 248108 {
		t.Errorf("Calculated solution (%d) was not expected", result)
	}
}
