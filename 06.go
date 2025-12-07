package main

import (
	"aoc_2025/lib"
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func scanFile(f string) *bufio.Scanner {
	file, _ := os.Open(f)
	return bufio.NewScanner(file)
}

func parseInput(scanner *bufio.Scanner) (lib.Map[int], []string) {
	var buf [][]int
	var ops string
	parsingNums := true

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '*' || line[0] == '+' {
			parsingNums = false
		}
		if parsingNums {
			buf = append(buf, intsPlease(line))
		} else {
			ops = scanner.Text()
			break
		}
	}
	return lib.Map[int]{Buf: buf, NumRows: len(buf), NumCols: len(buf[0])}, opsPlease(ops)
}

func intsPlease(s string) []int {
	return nonEmpties(s, func(s string) int { res, _ := strconv.Atoi(s); return res })
}

func opsPlease(s string) []string {
	return nonEmpties(s, func(s string) string { return s })
}

func nonEmpties[T any](s string, f func(string) T) []T {
	var ret []T
	for _, el := range strings.Split(s, " ") {
		if el != "" {
			ret = append(ret, f(el))
		}
	}
	return ret
}

func solve(terms []int, op string) int {
	ret := 0
	apply := func(a, b int) int { return a + b }
	if op == "*" {
		ret = 1
		apply = func(a, b int) int { return a * b }
	}
	for _, term := range terms {
		ret = apply(ret, term)
	}
	return ret
}

func main() {
	fName := "inputs/06.in"
	intMap, ops := parseInput(scanFile(fName))
	intMap.Transpose()

	p1 := 0
	for r, op := range ops {
		p1 += solve(intMap.GetRow(r), op)
	}

	println("p1:", p1)

	byteMap := lib.NewByteMap(scanFile(fName))
	byteMap.EqualizeRows(' ')
	byteMap.Transpose()

	p2, idx := 0, 0
	var terms []int
	for r := range byteMap.NumRows {
		row := strings.TrimFunc(string(byteMap.GetRow(r)), func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		if row == "" {
			p2 += solve(terms, ops[idx])
			terms = []int{}
			idx++
		} else {
			term, _ := strconv.Atoi(row)
			terms = append(terms, term)
		}
	}
	p2 += solve(terms, ops[idx])

	println("p2:", p2)
}
