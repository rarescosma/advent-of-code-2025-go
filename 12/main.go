package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("inputs/12.in")

	scanner := bufio.NewScanner(file)
	p1 := 0

	var acc = 0
	var shapes []int

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "x") {
			if strings.TrimSpace(line) == "" {
				shapes = append(shapes, acc)
				acc = 0
			} else {
				acc += strings.Count(line, "#")
			}
			continue
		}

		parts := strings.Split(line, ": ")
		hw := strings.Split(parts[0], "x")
		h, w := intPlease(hw[0]), intPlease(hw[1])
		var blocks []int
		for _, b := range strings.Split(parts[1], " ") {
			blocks = append(blocks, intPlease(b))
		}

		area := h * w
		bArea := 0
		for i, bNum := range blocks {
			bArea += bNum * shapes[i]
		}
		if bArea < area {
			p1 += 1
		}
	}
	println("p1:", p1)
}

func intPlease(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}
