package main

import (
	"aoc_2025/lib"
	"bufio"
	"os"
	"slices"
)

type NodeIdx = int

type Graph struct {
	m    *lib.Map[byte]
	memo map[NodeIdx]int
}

func newGraph(m *lib.Map[byte]) Graph {
	return Graph{
		m:    m,
		memo: make(map[NodeIdx]int),
	}
}

func (g *Graph) addNode(r, c int) int {
	idx := g.nodeIndex(r, c)

	for row := r - 1; row >= 0; row-- {
		upper := g.m.Get(row, c)
		if upper == 'S' {
			g.memo[idx] = 1
		} else if upper != '|' {
			break
		}
		// There's an implicit topological sort of splitters, so by the time we reach
		// the node with index 'idx' it's guaranteed that we've already calculated
		// the number of paths for its upper node.
		if c-1 >= 0 && g.m.Get(row, c-1) == '^' {
			g.memo[idx] += g.memo[g.nodeIndex(row, c-1)]
		}
		if c+1 < g.m.NumCols && g.m.Get(row, c+1) == '^' {
			g.memo[idx] += g.memo[g.nodeIndex(row, c+1)]
		}
	}

	return g.memo[idx]
}

func (g *Graph) nodeIndex(r, c int) NodeIdx {
	return r*g.m.NumCols + c
}

func main() {
	file, _ := os.Open("inputs/07.in")
	scanner := bufio.NewScanner(file)
	theMap := lib.NewByteMap(scanner)

	// Assumption: the first splitter is always on row 2, so set row 1 to '|' to hit it.
	sCol := slices.Index(theMap.GetRow(0), 'S')
	theMap.Set(1, sCol, '|')

	graph := newGraph(&theMap)
	graph.addNode(0, sCol)

	for row := 2; row < theMap.NumRows; row++ {
		for col := range theMap.NumCols {
			switch theMap.Get(row, col) {
			case '^':
				if theMap.Get(row-1, col) == '|' {
					graph.addNode(row, col)
					theMap.Set(row, col-1, '|')
					theMap.Set(row, col+1, '|')
				}
			case '.':
				if theMap.Get(row-1, col) == '|' {
					theMap.Set(row, col, '|')
				}
			}
		}
	}

	println("p1:", len(graph.memo))

	p2 := 0
	lastRow := theMap.NumRows - 1

	for col := range theMap.NumCols {
		if theMap.Get(lastRow, col) == '|' {
			p2 += graph.addNode(lastRow, col)
		}
	}

	println("p2:", p2)
}
