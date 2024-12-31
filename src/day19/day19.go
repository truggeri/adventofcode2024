package day19

import (
	"slices"
	"strings"
)

const THREAD_POOL = 32

type trie struct {
	chr  rune
	next map[rune]*trie
	end  bool
}

type thread struct {
	next    map[rune]*trie
	pattern string
}

func Solve(input string) uint {
	a, patterns := parseInput(input)
	t := buildPossibilities(a)
	var cnt uint = 0
	for _, p := range patterns {
		if isPatternPossible(t.next, p) {
			cnt++
		}
	}
	return cnt
}

func parseInput(input string) ([]string, []string) {
	lines := strings.Split(input, "\n")
	return strings.Split(lines[0], ", "), lines[2:]
}

func buildPossibilities(towels []string) trie {
	t := trie{next: make(map[rune]*trie)}
	for _, pattern := range towels {
		var prevT *trie = &t
		for i, r := range pattern {
			if _, ok := prevT.next[r]; !ok {
				tmp := trie{chr: r, next: make(map[rune]*trie)}
				prevT.next[r] = &tmp
			}
			prevT = prevT.next[r]
			if i == len(pattern)-1 {
				prevT.end = true
			}
		}
	}
	return t
}

func isPatternPossible(possibilities map[rune]*trie, pattern string) bool {
	threads := []thread{{next: possibilities, pattern: pattern}}
	coldStorage := make([]thread, 0)
	for {
		if len(threads) == 0 {
			if len(coldStorage) > 0 {
				x := THREAD_POOL
				if len(coldStorage) < x {
					x = len(coldStorage)
				}
				threads = coldStorage[:x]
				coldStorage = slices.Delete(coldStorage, 0, x)
			} else {
				return false
			}
		}

		next := make([]thread, 0)
		for i, t := range threads {
			if i > THREAD_POOL {
				coldStorage = append(coldStorage, t)
				continue
			}

			if v, ok := t.next[rune(t.pattern[0])]; ok {
				if len(t.pattern) == 1 {
					return true
				}

				thr := thread{pattern: t.pattern[1:], next: v.next}
				next = append([]thread{thr}, next...)
				if v.end {
					thr := thread{pattern: t.pattern[1:], next: possibilities}
					next = append([]thread{thr}, next...)
				}
			}
		}
		threads = next
	}
}
