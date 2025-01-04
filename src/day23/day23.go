package day23

import (
	"strings"
)

const FILTER_PREFIX = 't'

type connection struct {
	left, right string
}

func Solve(input string) uint {
	dict := make(map[string][]string)
	isConnected := make(map[connection]bool)
	connections := parseInput(input)
	for _, conn := range connections {
		if conn.left[0] == FILTER_PREFIX {
			dict[conn.left] = append(dict[conn.left], conn.right)
		}
		if conn.right[0] == FILTER_PREFIX {
			dict[conn.right] = append(dict[conn.right], conn.left)
		}
		isConnected[conn] = true
		isConnected[connection{left: conn.right, right: conn.left}] = true
	}

	var result, doubles, triples uint
	for _, v := range dict {
		if len(v) < 2 {
			continue
		}

		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				x := isConnected[connection{left: v[i], right: v[j]}]
				if x {
					result++
					if v[i][0] == FILTER_PREFIX || v[j][0] == FILTER_PREFIX {
						if v[i][0] == v[j][0] {
							triples++
						} else {
							doubles++
						}
					}
				}
			}
		}
	}
	return result - doubles/2 - 2*(triples/3)
}

func parseInput(input string) []connection {
	result := make([]connection, 0)
	for _, line := range strings.Split(input, "\n") {
		splits := strings.Split(line, "-")
		if len(splits) != 2 {
			continue
		}

		result = append(result, connection{left: splits[0], right: splits[1]})
	}
	return result
}
