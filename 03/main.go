package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("inputs/03.in")
	scanner := bufio.NewScanner(file)

	p1 := 0
	p2 := 0
	for scanner.Scan() {
		bat := scanner.Bytes()
		p1 += gobble(bat, 2)
		p2 += gobble(bat, 12)
	}

	println("p1:", p1)
	println("p2:", p2)
}

func digitIndexes(b []byte) [][]int {
	ret := make([][]int, 10)
	for i, c := range b {
		digit := int(c - '0')
		ret[digit] = append(ret[digit], i)
	}
	return ret
}

func gobble(b []byte, numNeeded int) int {
	indexes := digitIndexes(b)

	joltage := 0
	lastIdx := -1
	for numNeeded > 0 {
	out:
		for digit := 9; digit > 0; digit-- {
			for _, idx := range indexes[digit] {
				if idx <= len(b)-numNeeded && idx > lastIdx {
					joltage = joltage*10 + digit
					numNeeded--
					lastIdx = idx
					break out
				}
			}
		}
	}

	return joltage
}
