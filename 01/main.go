package main

import (
	"aoc_2025/lib"
	"bufio"
	"os"
)

type Dial struct {
	pos int
	max int
}

// SpinRight spins the dial in the "Right" direction.
//
// Returns the number of times it passed through 0
func (d *Dial) SpinRight(spin int) int {
	newPos, zeros := (d.pos+spin)%d.max, (d.pos+spin)/d.max
	d.pos = newPos
	return zeros
}

// SpinLeft Spins the dial in the "Left" direction.
//
// Returns the number of times it passed through 0
func (d *Dial) SpinLeft(spin int) int {
	newPos, zeros := (d.pos-spin)%d.max, (spin-d.pos)/d.max
	if newPos <= 0 && d.pos != 0 {
		zeros++
	}
	if newPos < 0 {
		newPos += d.max
	}
	d.pos = newPos
	return zeros
}

func main() {
	file, _ := os.Open("inputs/01.in")
	scanner := bufio.NewScanner(file)

	dial := Dial{pos: 50, max: 100}

	p1, p2 := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		spin := lib.IntPlease(line[1:])

		if direction == 'R' {
			p2 += dial.SpinRight(spin)
		} else {
			p2 += dial.SpinLeft(spin)
		}

		if dial.pos == 0 {
			p1 += 1
		}
	}
	println("p1:", p1)
	println("p2:", p2)
}
