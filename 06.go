package main

import (
	"aoc_2025/lib"
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func intMapFromScanner(scanner *bufio.Scanner) (lib.Map[int], []string) {
	var buf [][]int
	var ops string
	numRows, parsingNums := 0, true

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '*' || line[0] == '+' {
			parsingNums = false
		}
		if parsingNums {
			buf = append(buf, make([]int, 0))
			buf[numRows] = append(buf[numRows], intsPlease(line)...)
			numRows++
		} else {
			ops = scanner.Text()
			break
		}
	}
	numCols := len(buf[0])
	return lib.Map[int]{Buf: buf, NumRows: numRows, NumCols: numCols}, opsPlease(ops)
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

func intsPlease(s string) []int {
	return nonEmpties(s, func(s string) int {
		res, _ := strconv.Atoi(s)
		return res
	})
}

func opsPlease(s string) []string {
	return nonEmpties(s, func(s string) string {
		return s
	})
}

func scanFile(f string) *bufio.Scanner {
	file, _ := os.Open(f)
	return bufio.NewScanner(file)
}

func main() {
	fName := "inputs/06.in"
	var intMap lib.Map[int]
	intMap, ops := intMapFromScanner(scanFile(fName))
	intMap = intMap.Transpose()

	p1 := 0
	for r, op := range ops {
		part := 0
		if op == "*" {
			part = 1
			for _, el := range intMap.GetRow(r) {
				part = part * el
			}
		} else {
			for _, el := range intMap.GetRow(r) {
				part = part + el
			}
		}
		p1 += part
	}

	println("p1:", p1)

	var byteMap lib.Map[byte]
	byteMap = lib.NewByteMap(scanFile(fName))
	maxRow := 0
	for r := range byteMap.NumRows {
		maxRow = max(maxRow, len(byteMap.Buf[r]))
	}
	for r := range byteMap.NumRows {
		missing := maxRow - len(byteMap.Buf[r])
		for x := missing; x > 0; x-- {
			byteMap.Append(r, ' ')
		}
	}
	byteMap = byteMap.Transpose()

	idx := 0
	op := ops[idx]
	part := 0
	if op == "*" {
		part = 1
	}
	p2 := 0
	for r := range byteMap.NumRows {
		row := strings.TrimFunc(string(byteMap.GetRow(r)[:]), func(r rune) bool {
			return r == '*' || r == '+' || unicode.IsSpace(r)
		})
		if row != "" {
			term, _ := strconv.Atoi(row)
			if op == "*" {
				part = part * term
			} else {
				part = part + term
			}
		} else {
			p2 += part
			idx++
			op = ops[idx]
			if op == "*" {
				part = 1
			} else {
				part = 0
			}
		}
	}
	p2 += part
	println("p2:", p2)
}
