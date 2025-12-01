package main

import (
	"bufio"
	"os"
	"strconv"
)

const MODULUS = 100

// Spins the wheel in the "Right" direction.
//
// Returns the new value of the wheel and the number of times it passed through 0
func youSpinMeRightRound(wheel int, spin int) (int, int) {
	return (wheel + spin) % MODULUS, (wheel + spin) / MODULUS
}

// Spins the wheel in the "Left" direction.
//
// Returns the new value of the wheel and the number of times it passed through 0
func youSpinMeLeftRound(wheel int, spin int) (int, int) {
	newWheel, zeros := (wheel-spin)%MODULUS, (spin-wheel)/MODULUS
	if newWheel <= 0 && wheel != 0 {
		zeros++
	}
	if newWheel < 0 {
		newWheel += MODULUS
	}
	return newWheel, zeros
}

func main() {
	file, _ := os.Open("inputs/01.in")
	scanner := bufio.NewScanner(file)

	wheel, p1, p2, zeros := 50, 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		spin, _ := strconv.Atoi(line[1:])

		if direction == 'R' {
			wheel, zeros = youSpinMeRightRound(wheel, spin)
		} else {
			wheel, zeros = youSpinMeLeftRound(wheel, spin)
		}

		p2 += zeros
		if wheel == 0 {
			p1 += 1
		}
	}
	println("p1:", p1)
	println("p2:", p2)
}
