package main

import (
	"bufio"
	"os"
)

type Pos struct {
	row int
	col int
}

type Map struct {
	buf     [][]byte
	numRows int
	numCols int
}

func (m *Map) Get(p Pos) byte {
	return m.buf[p.row][p.col]
}

func (m *Map) Set(p Pos, b byte) {
	m.buf[p.row][p.col] = b
}

func MapFromScanner(scanner *bufio.Scanner) Map {
	var buf [][]byte
	numRows := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		buf = append(buf, make([]byte, 0))
		buf[numRows] = append(buf[numRows], line...)
		numRows++
	}
	numCols := len(buf[0])
	return Map{buf, numRows, numCols}
}

func main() {
	file, _ := os.Open("inputs/04.in")
	scanner := bufio.NewScanner(file)

	theMap := MapFromScanner(scanner)

	p1 := removeRolls(&theMap)
	p2, removed := p1, p1

	for removed > 0 {
		removed = removeRolls(&theMap)
		p2 += removed
	}
	println("p1:", p1)
	println("p2:", p2)
}

func removeRolls(m *Map) int {
	ret := 0
	var toRemove []Pos
	for x := 0; x < m.numCols; x++ {
		for y := 0; y < m.numRows; y++ {
			if m.Get(Pos{x, y}) != '@' {
				continue
			}
			neighs := 0
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue
					}
					nx, ny := x+dx, y+dy
					if m.numRows > nx && nx >= 0 && m.numCols > ny && ny >= 0 && m.Get(Pos{nx, ny}) == '@' {
						neighs += 1
					}
				}
			}
			if neighs < 4 {
				toRemove = append(toRemove, Pos{x, y})
				ret += 1
			}
		}
	}
	for _, pos := range toRemove {
		m.Set(pos, 'x')
	}
	return ret
}
