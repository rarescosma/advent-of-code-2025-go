package main

import (
	"bufio"
	"os"
	"strings"
)

func numPathsCached(
	cache *map[string]int,
	rev *map[string][]string,
	from, to string,
	exclude *map[string]bool,
) int {
	if to == from {
		return 1
	}

	if ret, ok := (*cache)[to]; ok {
		return ret
	}

	ret := 0
	for _, prev := range (*rev)[to] {
		if (*exclude)[prev] {
			continue
		}
		ret += numPathsCached(cache, rev, from, prev, exclude)
	}

	(*cache)[to] = ret
	return ret
}

type Graph struct {
	rev map[string][]string
}

func (g *Graph) numPaths(from, to string, exclude []string) int {
	cache := make(map[string]int)
	var excludeSet = make(map[string]bool, len(exclude))
	for _, el := range exclude {
		excludeSet[el] = true
	}
	return numPathsCached(&cache, &g.rev, from, to, &excludeSet)
}

func main() {
	file, _ := os.Open("inputs/11.in")

	g := Graph{make(map[string][]string)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		from := parts[0]
		for _, to := range strings.Split(parts[1], " ") {
			g.rev[to] = append(g.rev[to], from)
		}
	}

	p1 := g.numPaths("you", "out", []string{})
	println("p1:", p1)

	var p2 int
	fd := g.numPaths("fft", "dac", []string{"out", "svr"})
	if fd == 0 {
		p2 = g.numPaths("svr", "dac", []string{"out", "fft"}) *
			g.numPaths("dac", "fft", []string{"out", "svr"}) *
			g.numPaths("fft", "out", []string{"dac", "svr"})
	} else {
		p2 = fd * g.numPaths("svr", "fft", []string{"out", "dac"}) *
			g.numPaths("dac", "out", []string{"fft", "svr"})
	}

	println("p2:", p2)
}
