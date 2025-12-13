package main

import (
	"bufio"
	"os"
	"strings"
)

func numPathsCached(
	cache *map[string]int,
	rev *map[string][]string,
	from string,
	to string,
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

func main() {
	file, _ := os.Open("inputs/11.in")

	rev := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		from := parts[0]
		for _, to := range strings.Split(parts[1], " ") {
			rev[to] = append(rev[to], from)
		}
	}

	var cache = make(map[string]int)
	p1 := numPathsCached(&cache, &rev, "you", "out", &map[string]bool{})
	println("p1:", p1)

	cache = make(map[string]int)
	sd := numPathsCached(&cache, &rev, "svr", "dac", &map[string]bool{"out": true, "fft": true})
	cache = make(map[string]int)
	df := numPathsCached(&cache, &rev, "dac", "fft", &map[string]bool{"out": true, "svr": true})
	cache = make(map[string]int)
	fo := numPathsCached(&cache, &rev, "fft", "out", &map[string]bool{"dac": true, "svr": true})

	cache = make(map[string]int)
	sf := numPathsCached(&cache, &rev, "svr", "fft", &map[string]bool{"out": true, "dac": true})
	cache = make(map[string]int)
	fd := numPathsCached(&cache, &rev, "fft", "dac", &map[string]bool{"out": true, "svr": true})
	cache = make(map[string]int)
	do := numPathsCached(&cache, &rev, "dac", "out", &map[string]bool{"fft": true, "svr": true})

	p2 := sd*df*fo + sf*fd*do
	println("p2:", p2)
}
