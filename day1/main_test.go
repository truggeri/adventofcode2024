package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	a := []uint{3, 4, 2, 1, 3, 3}
	b := []uint{4, 3, 5, 3, 9, 3}
	result := Solve(a, b)
	if result != 11 {
		t.Errorf("Calculated solution was not expected")
	}
}
