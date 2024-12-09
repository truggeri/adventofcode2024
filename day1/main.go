package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(Solve([]uint{3, 4, 2, 1, 3, 3}, []uint{4, 3, 5, 3, 9, 3}))
}

func Solve(a, b []uint) uint {
	slices.Sort(a)
	slices.Sort(b)
	var sum uint = 0
	for i := range a {
		sum = sum + absDiffUint(a[i], b[i])
	}
	return sum
}

// @see {https://stackoverflow.com/a/68277627}
func absDiffUint(x, y uint) uint {
	if x < y {
		return y - x
	}
	return x - y
}
