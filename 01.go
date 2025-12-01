package main

import (
	"bufio"
	"os"
	"strconv"
)

const MODULUS = 100

func zeros(wheel int, antiWheel int, spin int) (int, int, int) {
	newWheel := wheel + spin
	return newWheel % MODULUS, (antiWheel + MODULUS - (spin % MODULUS)) % MODULUS, newWheel / MODULUS
}

func main() {
	file, _ := os.Open("inputs/01.in")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	scanner := bufio.NewScanner(file)

	/*
	   The math is much simpler for the positive case, so just
	   spin two wheels in reverse directions:
	   - wR spins R on R, L on L
	   - wL spins L on R, R on L

	   When spinning right, count 0 passes for wR only.
	   When spinning left, count 0 passes for wL only.
	   Sum the contributions.
	*/
	wR, wL, p1, p2, rightZeros, leftZeros := 50, 50, 0, 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		spin, _ := strconv.Atoi(line[1:])

		if direction == "R" {
			wR, wL, rightZeros = zeros(wR, wL, spin)
			p2 += rightZeros
		} else {
			wL, wR, leftZeros = zeros(wL, wR, spin)
			p2 += leftZeros
		}
		if wR == 0 {
			p1 += 1
		}
	}
	println("p1:", p1)
	println("p2:", p2)
}
