package main

import (
	"aoc_2025/lib"
	"bufio"
	"math"
	"os"
	"runtime"
	"strings"
)

type State []int
type Button []int
type ButtonIdx = int
type Combo []ButtonIdx

type Problem struct {
	p1Goal, p2Goal State
	buttons        []Button
}

type Ans struct{ p1, p2 int }

func main() {
	file, _ := os.Open("inputs/10.in")

	scanner := bufio.NewScanner(file)

	var problems []Problem

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line[1:len(line)-1], "] (")

		p1Goal := make(State, len(parts[0]))
		for i, c := range parts[0] {
			if c == '#' {
				p1Goal[i] = 1
			}
		}

		parts = strings.Split(parts[1], ") {")
		p2Goal := lib.IntsPlease(parts[1], ",")

		buttonParts := strings.Split(parts[0], ") (")
		buttons := make([]Button, len(buttonParts))
		for i, s := range buttonParts {
			buttons[i] = lib.IntsPlease(s, ",")
		}
		problems = append(problems, Problem{p1Goal, p2Goal, buttons})
	}

	pool := lib.NewPool(runtime.NumCPU(), problems, func(chunk []Problem) Ans {
		p1, p2 := 0, 0
		for _, problem := range chunk {
			parityMap := make(map[uint64][]Combo)
			p1 += p1Solve(problem.p1Goal, problem.buttons, &parityMap)

			cache := make(map[uint64]int)
			cache[hashState(make(State, len(problem.p2Goal)))] = 0
			p2 += p2Solve(problem.p2Goal, problem.buttons, parityMap, cache)
		}
		return Ans{p1, p2}
	})

	p1, p2 := 0, 0
	for ans := range pool.Go() {
		p1 += ans.p1
		p2 += ans.p2
	}
	println("p1:", p1)
	println("p2:", p2)
}

func p1Solve(goal State, buttons []Button, parityMap *map[uint64][]Combo) int {
	nButtons := len(buttons)
	limit := 1 << nButtons
	ret := math.MaxInt

	currentParity := make(State, len(goal))

	for i := 0; i < limit; i++ {
		for k := range currentParity {
			currentParity[k] = 0
		}

		var bits Combo

		for bit := 0; bit < nButtons; bit++ {
			if (i & (1 << bit)) != 0 {
				bits = append(bits, bit)
				for _, pos := range buttons[bit] {
					currentParity[pos] ^= 1
				}
			}
		}

		numBits := len(bits)
		if numBits < ret && stateEq(currentParity, goal) {
			ret = numBits
		}

		key := hashState(currentParity)
		(*parityMap)[key] = append((*parityMap)[key], bits)
	}
	return ret
}

func p2Solve(goal State, buttons []Button, parityMap map[uint64][]Combo, cache map[uint64]int) int {
	key := hashState(goal)
	if res, ok := cache[key]; ok {
		return res
	}

	ret := math.MaxInt

	nextGoal := make(State, len(goal))
out:
	for _, bits := range parityMap[hashParity(goal)] {
		copy(nextGoal, goal)

		for _, bit := range bits {
			for _, pos := range buttons[bit] {
				nextGoal[pos]--
			}
		}

		for i, v := range nextGoal {
			if v < 0 {
				continue out
			}
			nextGoal[i] = v / 2
		}

		if solve := p2Solve(nextGoal, buttons, parityMap, cache); solve != math.MaxInt {
			total := 2*solve + len(bits)
			if total < ret {
				ret = total
			}
		}
	}

	cache[key] = ret
	return ret
}

func hashState(s State) uint64 {
	return fastHash(s, func(v int) uint64 {
		return uint64(v)
	})
}

func hashParity(s State) uint64 {
	return fastHash(s, func(v int) uint64 {
		return uint64(v % 2)
	})
}

func fastHash(s []int, fn func(v int) uint64) uint64 {
	var h uint64 = 14695981039346656037
	for _, v := range s {
		h ^= fn(v)
		h *= 1099511628211
	}
	return h
}

func stateEq(s1, s2 State) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
