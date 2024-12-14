package day3

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result := Solve(input)
	if result != 161 {
		t.Errorf("Calculated solution was not expected")
	}
}
