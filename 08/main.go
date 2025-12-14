package main

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Junc struct {
	x, y, z int
}

type Edge struct {
	u, v int
	dist int
}

type DSU struct {
	parent []int
	size   []int
	count  int // number of disjoint sets
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent: parent, size: size, count: n}
}

func (d *DSU) Find(i int) int {
	root := i
	for d.parent[root] != root {
		root = d.parent[root]
	}

	// path compression
	curr := i
	for curr != root {
		next := d.parent[curr]
		d.parent[curr] = root
		curr = next
	}

	return root
}

func (d *DSU) Union(i, j int) bool {
	rootI := d.Find(i)
	rootJ := d.Find(j)

	if rootI == rootJ {
		return false
	}

	// union by size: attach smaller tree to larger tree
	if d.size[rootI] < d.size[rootJ] {
		rootI, rootJ = rootJ, rootI
	}

	d.parent[rootJ] = rootI
	d.size[rootI] += d.size[rootJ]
	d.count--
	return true
}

func main() {
	takeDists := 1000
	file, _ := os.Open("inputs/08.in")

	var juncs []Junc
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		juncs = append(juncs, getJunc(scanner.Text()))
	}
	n := len(juncs)

	edges := make([]Edge, 0, n*(n-1)/2)

	for i := 0; i < n; i++ {
		p1 := juncs[i]
		for j := i + 1; j < n; j++ {
			p2 := juncs[j]
			edges = append(edges, Edge{u: i, v: j, dist: dist(p1, p2)})
		}
	}

	slices.SortFunc(edges, func(a, b Edge) int {
		return a.dist - b.dist
	})

	dsu := NewDSU(n)
	p1, p2 := 1, 0
	numDists := 0

	totalEdges := len(edges)
	i := 0

	for i < totalEdges {
		currentDist := edges[i].dist
		j := i

		// process the "bucket" of all edges with this specific distance
		for j < totalEdges && edges[j].dist == currentDist {
			e := edges[j]
			if dsu.Union(e.u, e.v) && dsu.count == 1 {
				p2 = juncs[e.u].x * juncs[e.v].x
				if numDists >= takeDists {
					goto done
				}
			}
			j++
		}

		numDists++

		// p1 logic: triggered exactly when we hit takeDists unique distances
		if numDists == takeDists {
			var componentSizes []int
			for k := 0; k < n; k++ {
				if dsu.parent[k] == k {
					componentSizes = append(componentSizes, dsu.size[k])
				}
			}

			sort.Ints(componentSizes)
			count := len(componentSizes)

			if count >= 3 {
				p1 = componentSizes[count-1] * componentSizes[count-2] * componentSizes[count-3]
			}

			if dsu.count == 1 {
				goto done
			}
		}

		i = j
	}

done:
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

func dist(j1, j2 Junc) int {
	return (j1.x-j2.x)*(j1.x-j2.x) + (j1.y-j2.y)*(j1.y-j2.y) + (j1.z-j2.z)*(j1.z-j2.z)
}
