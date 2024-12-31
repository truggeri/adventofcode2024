package day17

import (
	"strings"
	"testing"
)

func compareOutBuffer(t *testing.T, given, expected []string) {
	if len(given) != len(expected) {
		t.Errorf("Given and expected string slices not of equal length")
	}
	for i := range given {
		if given[i] != expected[i] {
			t.Errorf("Given string slice did not match expected")
		}
	}
}

func TestSmall1(t *testing.T) {
	cmp := computer{regC: 9}
	cmp.instructions = append(cmp.instructions, 2, 6)
	cmp.tick()
	if cmp.regB != 1 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestSmall2(t *testing.T) {
	cmp := computer{regA: 10}
	cmp.instructions = append(cmp.instructions, 5, 0, 5, 1, 5, 4)
	for {
		if cmp.tick() {
			break
		}
	}
	compareOutBuffer(t, cmp.outBuffer, strings.Split("0,1,2", ","))
}

func TestSmall3(t *testing.T) {
	cmp := computer{regA: 2024}
	cmp.instructions = append(cmp.instructions, 0, 1, 5, 4, 3, 0)
	for {
		if cmp.tick() {
			break
		}
	}
	if cmp.regA != 0 {
		t.Errorf("Calculated solution was not expected")
	}
	compareOutBuffer(t, cmp.outBuffer, strings.Split("4,2,5,6,7,7,7,7,3,1,0", ","))

}

func TestSmall4(t *testing.T) {
	cmp := computer{regB: 29}
	cmp.instructions = append(cmp.instructions, 1, 7)
	cmp.tick()
	if cmp.regB != 26 {
		t.Errorf("Calculated solution was not expected")
	}
}

func TestSmall5(t *testing.T) {
	cmp := computer{regB: 2024, regC: 43690}
	cmp.instructions = append(cmp.instructions, 4, 0)
	cmp.tick()
	if cmp.regB != 44354 {
		t.Errorf("Calculated solution was not expected")
	}
}
