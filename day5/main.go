package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type pageNumber uint
type rule pageNumber
type batch []pageNumber

func main() {
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
	fmt.Println("Solution:", Solve(input))
}

func Solve(input string) uint {
	rules, batches := parseInput(input)
	var count uint = 0
	for _, b := range batches {
		count += uint(medianMatchingPage(rules, b))
	}
	return count
}

func parseInput(input string) (map[pageNumber][]rule, []batch) {
	rules := make(map[pageNumber][]rule)
	batches := make([]batch, 0)
	parsingRules := true
	b := 0

	for _, i := range strings.Split(input, "\n") {
		if i == "" {
			parsingRules = false
		} else if parsingRules {
			ruleValues := strings.Split(i, "|")
			if len(ruleValues) != 2 {
				continue
			}

			x, err := strconv.Atoi(ruleValues[0])
			if err != nil {
				continue
			}
			y, err := strconv.Atoi(ruleValues[1])
			if err != nil {
				continue
			}
			rules[pageNumber(x)] = slices.Insert(rules[pageNumber(x)], len(rules[pageNumber(x)]), rule(y))
		} else {
			pageNumbers := strings.Split(i, ",")
			batches = slices.Insert(batches, b, make(batch, 0))
			for _, num := range pageNumbers {
				x, err := strconv.Atoi(num)
				if err != nil {
					continue
				}
				batches[b] = slices.Insert(batches[b], len(batches[b]), pageNumber(x))
			}
			b += 1
		}
	}

	return rules, batches
}

func medianMatchingPage(rules map[pageNumber][]rule, b batch) pageNumber {
	if !isBatchValid(rules, b) {
		return 0
	}

	return b[len(b)/2]
}

func isBatchValid(rules map[pageNumber][]rule, b batch) bool {
	pagesPrintedAlready := make(map[pageNumber]bool)
	for _, p := range b {
		for _, r := range rules[p] {
			if pagesPrintedAlready[pageNumber(r)] {
				return false
			}
		}
		pagesPrintedAlready[p] = true
	}
	return true
}
