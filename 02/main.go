package main

import (
	"aoc_2025/lib"
	"os"
	"strings"
	"sync"
)

// Precomputed powers of 10 to avoid math.Pow
var powersOf10 = [...]int{
	1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000,
	10000000000, 100000000000, 1000000000000, 10000000000000, 100000000000000,
	1000000000000000, 10000000000000000, 100000000000000000, 1000000000000000000,
}

type Ans struct {
	p1 int
	p2 int
}

func main() {
	file, _ := os.ReadFile("inputs/02.in")
	ranges := strings.Split(strings.TrimSpace(string(file)), ",")

	var wg sync.WaitGroup
	wg.Add(len(ranges))
	ansChan := make(chan Ans)

	go func() {
		wg.Wait()
		close(ansChan)
	}()

	for _, rng := range ranges {
		go func(rng string, wg *sync.WaitGroup) {
			defer wg.Done()

			p1, p2 := 0, 0
			ints := lib.IntsPlease(rng, "-")
			beg, end := ints[0], ints[1]

			for x := beg; x <= end; x++ {
				numDigits, p2Found := fastDigits(x), false

				for rad := 1; rad < numDigits; rad++ {
					if numDigits%rad == 0 && numDigits/rad >= 2 {
						if isInvalid(x, rad) {
							if numDigits/rad == 2 {
								p1 += x
							}
							if !p2Found {
								p2 += x
								p2Found = true
							}
						}
					}
				}
			}
			ansChan <- Ans{p1: p1, p2: p2}
		}(rng, &wg)
	}

	p1, p2 := 0, 0
	for result := range ansChan {
		p1 += result.p1
		p2 += result.p2
	}
	println("p1:", p1)
	println("p2:", p2)
}

// Fast integer-only digit counting
func fastDigits(x int) int {
	for i, pow := range powersOf10 {
		if x < pow {
			return i
		}
	}
	return len(powersOf10)
}

func isInvalid(x int, radius int) bool {
	base := powersOf10[radius]
	rem := x % base
	x = x / base

	for x > base {
		oldRem := rem
		rem = x % base
		if rem != oldRem {
			return false
		}
		x = x / base
	}

	return x == rem
}
