package day4

import (
	"slices"
	"strings"
	"unicode/utf8"
)

const WORD_TO_FIND = "XMAS"

var _rows []string

func Solve(input string) uint {
	_rows = make([]string, 0)
	return matchesByStrategy(input, rows) + matchesByStrategy(input, columns) + matchesByStrategy(input, diagonals)
}

func matchesByStrategy(input string, strategy func(string) []string) uint {
	var count uint = 0
	for _, x := range strategy(input) {
		if contains(x) {
			count += 1
		}
		if contains(reverse(x)) {
			count += 1
		}
	}
	return count
}

func contains(input string) bool {
	return strings.Contains(input, WORD_TO_FIND)
}

func rows(input string) []string {
	if len(_rows) == 0 {
		_rows = strings.Split(input, "\n")
	}
	return _rows
}

func columns(input string) []string {
	rows := rows(input)
	var cols []string = make([]string, len(rows))
	for _, r := range rows {
		for j, v := range r {
			if j >= len(rows) {
				panic("The input puzzle is not square")
			}
			cols[j] += string(v)
		}
	}
	return cols
}

func diagonals(input string) []string {
	rows := rows(input)
	wordSize := len(WORD_TO_FIND)
	squareSize := len(rows)
	var diagonals []string = make([]string, 0)
	seed := squareSize - wordSize

	for s := 0; s <= seed; s++ {
		buf1 := forwardLower(s, squareSize, rows)
		buf2 := forwardUpper(s, squareSize, rows)
		buf3 := backLower(s, squareSize, rows)
		buf4 := backUpper(s, squareSize, rows)
		diagonals = slices.Insert(diagonals, len(diagonals), string(buf1), string(buf2), string(buf3), string(buf4))
	}

	return diagonals
}

func forwardLower(s int, squareSize int, rows []string) []byte {
	buf := make([]byte, 0)
	i := 0
	j := s
	for {
		if !inPuzzle(i, j, squareSize) {
			break
		}

		buf = slices.Insert(buf, len(buf), rows[i][j])
		i++
		j++
	}
	return buf
}

func forwardUpper(s int, squareSize int, rows []string) []byte {
	buf := make([]byte, 0)
	if s == 0 {
		return buf
	}
	i := s
	j := 0
	for {
		if !inPuzzle(i, j, squareSize) {
			break
		}

		buf = slices.Insert(buf, len(buf), rows[i][j])
		i++
		j++
	}
	return buf
}

func backUpper(s int, squareSize int, rows []string) []byte {
	buf := make([]byte, 0)
	i := squareSize - 1 - s
	j := 0
	for {
		if !inPuzzle(i, j, squareSize) {
			break
		}

		buf = slices.Insert(buf, len(buf), rows[i][j])
		i--
		j++
	}
	return buf
}

func backLower(s int, squareSize int, rows []string) []byte {
	buf := make([]byte, 0)
	i := squareSize - 1
	j := s
	for {
		if !inPuzzle(i, j, squareSize) {
			break
		}

		buf = slices.Insert(buf, len(buf), rows[i][j])
		i--
		j++
	}
	return buf
}

// @see {https://stackoverflow.com/a/34521190}
func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func inPuzzle(i, j int, squareSize int) bool {
	return i >= 0 && i < squareSize && j >= 0 && j < squareSize
}
