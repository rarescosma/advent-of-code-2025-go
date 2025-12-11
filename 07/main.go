package main

import (
	"aoc_2025/lib"
	"bufio"
	"os"
	"slices"
)

type NodeIdx = int

type Graph struct {
	m        *lib.Map[byte]
	numPaths map[NodeIdx]int
}

func newGraph(m *lib.Map[byte]) Graph {
	return Graph{m, make(map[NodeIdx]int)}
}

func (g *Graph) addNode(r, c int) int {
	idx := g.nodeIndex(r, c)

	for rowAbove := r - 1; rowAbove >= 0; rowAbove-- {
		upper := g.m.Get(rowAbove, c)
		// Walk up the beam.
		if upper != '|' {
			// If we hit the source we must be the first splitter, so there's only 1 path.
			if upper == 'S' {
				g.numPaths[idx] = 1
			}
			break
		}

		// Splitters are added to the graph from top to bottom (with increasing row numbers),
		// which generates an implicit topological ordering.
		//
		// It's therefore guaranteed that we've already calculated the number of paths
		// to all the splitters above the current splitter.
		if c-1 >= 0 && g.m.Get(rowAbove, c-1) == '^' {
			g.numPaths[idx] += g.numPaths[g.nodeIndex(rowAbove, c-1)]
		}
		if c+1 < g.m.NumCols && g.m.Get(rowAbove, c+1) == '^' {
			g.numPaths[idx] += g.numPaths[g.nodeIndex(rowAbove, c+1)]
		}
	}

	return g.numPaths[idx]
}

func (g *Graph) nodeIndex(r, c int) NodeIdx {
	return r*g.m.NumCols + c
}

func main() {
	file, _ := os.Open("inputs/07.in")
	scanner := bufio.NewScanner(file)
	theMap := lib.NewByteMap(scanner)
	graph := newGraph(&theMap)

	// Assumption: the first splitter is always directly below the source on row 2,
	// so set row 1 to '|' to hit it.
	theMap.Set(1, slices.Index(theMap.GetRow(0), 'S'), '|')

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

	println("p1:", len(graph.numPaths))

	p2 := 0
	lastRow := theMap.NumRows - 1

	for col := range theMap.NumCols {
		if theMap.Get(lastRow, col) == '|' {
			p2 += graph.addNode(lastRow, col)
		}
	}

	println("p2:", p2)
}
