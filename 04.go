package main

import (
	"aoc_2025/lib"
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("inputs/04.in")
	scanner := bufio.NewScanner(file)

	theMap := lib.NewByteMap(scanner)

	p1 := removeRolls(&theMap)
	p2, removed := p1, p1

	for removed > 0 {
		removed = removeRolls(&theMap)
		p2 += removed
	}
	println("p1:", p1)
	println("p2:", p2)
}

func removeRolls(m *lib.Map[byte]) int {
	ret := 0
	var toRemove []lib.Pos
	for x := 0; x < m.NumCols; x++ {
		for y := 0; y < m.NumRows; y++ {
			if m.Get(lib.Pos{Row: x, Col: y}) != '@' {
				continue
			}
			neighs := 0
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue
					}
					nx, ny := x+dx, y+dy
					if m.NumRows > nx && nx >= 0 && m.NumCols > ny && ny >= 0 && m.Get(lib.Pos{Row: nx, Col: ny}) == '@' {
						neighs += 1
					}
				}
			}
			if neighs < 4 {
				toRemove = append(toRemove, lib.Pos{Row: x, Col: y})
				ret += 1
			}
		}
	}
	for _, pos := range toRemove {
		m.Set(pos, 'x')
	}
	return ret
}
