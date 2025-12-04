package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("inputs/04.in")
	scanner := bufio.NewScanner(file)

	var theMap [][]byte
	numRows := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		theMap = append(theMap, make([]byte, 0))
		theMap[numRows] = append(theMap[numRows], line...)
		numRows++
	}
	numCols := len(theMap[0])

	p1 := removeRolls(&theMap, numRows, numCols)
	p2, removed := p1, p1

	for removed > 0 {
		removed = removeRolls(&theMap, numRows, numCols)
		p2 += removed
	}
	println("p1:", p1)
	println("p2:", p2)
}

func removeRolls(m *[][]byte, numRows int, numCols int) int {
	ret := 0
	var toRemove [][2]int
	for x := 0; x < numCols; x++ {
		for y := 0; y < numRows; y++ {
			if (*m)[x][y] != '@' {
				continue
			}
			neighs := 0
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue
					}
					nx, ny := x+dx, y+dy
					if numRows > nx && nx >= 0 && numCols > ny && ny >= 0 && (*m)[nx][ny] == '@' {
						neighs += 1
					}
				}
			}
			if neighs < 4 {
				toRemove = append(toRemove, [2]int{x, y})
				ret += 1
			}
		}
	}
	for _, xy := range toRemove {
		(*m)[xy[0]][xy[1]] = 'x'
	}
	return ret
}
