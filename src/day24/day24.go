package day24

import (
	"strconv"
	"strings"
)

const OUTPUT_WIRE_PREFIX = 'z'

type gate struct {
	typ         string
	left, right string
}

func (g gate) eval(left, right uint8) uint8 {
	switch g.typ {
	case "AND":
		return left & right
	case "OR":
		return left | right
	case "XOR":
		return left ^ right
	default:
		panic("Gate type could not be determined")
	}
}

func Solve(input string) uint {
	wires, gates := parseInput(input)
	for {
		if len(gates) == 0 {
			break
		}

		updates := 0
		for k, g := range gates {
			left, okL := wires[g.left]
			right, okR := wires[g.right]
			if okL && okR {
				wires[k] = g.eval(left, right)
				delete(gates, k)
				updates++
			}
		}

		if updates == 0 {
			panic("Iterated through gates and could not make forward progress")
		}
	}
	return outputWire(wires)
}

func parseInput(input string) (map[string]uint8, map[string]gate) {
	wires := make(map[string]uint8)
	gates := make(map[string]gate)
	parseGates := false
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			parseGates = true
			continue
		}

		if parseGates {
			d := strings.Split(line, " ")
			if len(d) != 5 || d[3] != "->" {
				continue
			}

			gates[d[4]] = gate{typ: d[1], left: d[0], right: d[2]}
		} else {
			d := strings.Split(line, ": ")
			if len(d) != 2 {
				continue
			}

			var v uint8
			if d[1] == "1" {
				v = 1
			}
			wires[d[0]] = v
		}
	}
	return wires, gates
}

func outputWire(wires map[string]uint8) uint {
	var result uint
	for k, w := range wires {
		if k[0] != OUTPUT_WIRE_PREFIX {
			continue
		}

		pos, _ := strconv.Atoi(k[1:])
		result = result | uint(w)<<pos
	}
	return result
}
