package main

import (
	"bufio"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Junc struct {
	x, y, z int
}

type JuncPair struct {
	left, right JuncIdx
}

type JuncIdx = int
type JuncDist = int
type DistMap = map[JuncDist][]JuncPair

type CircuitIdx = int

func main() {
	takeDists := 1000
	file, _ := os.Open("inputs/08.in")
	scanner := bufio.NewScanner(file)

	var juncs []Junc
	juncIdx := 0
	distMap := make(DistMap)
	for scanner.Scan() {
		junc := getJunc(scanner.Text())
		juncs = append(juncs, junc)
		for other := range juncIdx {
			key := dist(junc, juncs[other])
			distMap[key] = append(distMap[key], JuncPair{juncIdx, other})
		}
		juncIdx++
	}

	var circuits []CircuitIdx
	for juncIdx := range juncs {
		circuits = append(circuits, juncIdx)
	}

	p1, p2, numDists, numCircuits := 1, 0, 0, len(circuits)
out:
	for _, dist := range slices.Sorted(maps.Keys(distMap)) {
		for _, pair := range distMap[dist] {
			c1, c2 := circuits[pair.left], circuits[pair.right]
			if c1 == c2 {
				continue
			}

			for k, v := range circuits {
				if v == c2 {
					circuits[k] = c1
				}
			}
			numCircuits--
			if numCircuits == 1 {
				p2 = juncs[pair.left].x * juncs[pair.right].x
				if numDists > takeDists {
					break out
				}
			}
		}
		numDists++

		if numDists == takeDists {
			tally := make(map[CircuitIdx]int)
			for _, circuitIdx := range circuits {
				tally[circuitIdx]++
			}
			for _, circuitSize := range slices.Sorted(maps.Values(tally))[len(tally)-3:] {
				p1 *= circuitSize
			}
		}
	}

	println("p1:", p1)
	println("p2:", p2)
}

func getJunc(s string) Junc {
	var coords []int
	for _, el := range strings.Split(s, ",") {
		if el != "" {
			coord, _ := strconv.Atoi(el)
			coords = append(coords, coord)
		}
	}
	return Junc{coords[0], coords[1], coords[2]}
}

func dist(j1, j2 Junc) JuncDist {
	return (j1.x-j2.x)*(j1.x-j2.x) + (j1.y-j2.y)*(j1.y-j2.y) + (j1.z-j2.z)*(j1.z-j2.z)
}
