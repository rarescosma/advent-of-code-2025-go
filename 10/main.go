package main

import (
	"bufio"
	"encoding/binary"
	"hash/fnv"
	"math"
	"os"
	"strconv"
	"strings"
)

type State = []int

type Move = []int

func main() {
	file, _ := os.Open("inputs/10.in")

	scanner := bufio.NewScanner(file)

	p1, p2 := 0, 0
	for scanner.Scan() {
		line := scanner.Text()[1:]
		line = line[:len(line)-1]
		parts := strings.Split(line, "] (")
		goal := make([]int, len(parts[0]))
		for i, c := range parts[0] {
			if c == '#' {
				goal[i] = 1
			}
		}

		parts = strings.Split(parts[1], ") {")

		var p2Goal []int
		for _, val := range intsPlease(parts[1]) {
			p2Goal = append(p2Goal, val)
		}

		var moves []Move
		for _, part := range strings.Split(parts[0], ") (") {
			moves = append(moves, intsPlease(part))
		}
		pCache := make(map[uint32][][]Move)
		minMoves := math.MaxInt
		for _, combo := range patternsFor(goal, moves, &pCache) {
			if len(combo) < minMoves {
				minMoves = len(combo)
			}
		}
		p1 += minMoves
		cache := make(map[uint32]int)
		cache[stateHash(make(State, len(goal)))] = 0
		p2 += p2Solve(p2Goal, moves, &cache, &pCache)
	}

	println("p1:", p1)
	println("p2:", p2)
}

func p2Solve(goal []int, moves []Move, cache *map[uint32]int, pCache *map[uint32][][]Move) int {
	key := stateHash(goal)
	if res, ok := (*cache)[key]; ok {
		return res
	}

	state := p2Goal2State(&goal)
	combos := patternsFor(state, moves, pCache)

	var rec []int
out:
	for _, combo := range combos {
		newGoal := make(State, len(goal))
		copy(newGoal, goal)
		for _, move := range combo {
			p2Transform(&newGoal, move)
		}
		for _, digit := range newGoal {
			if digit < 0 {
				continue out
			}
		}
		for i := range newGoal {
			newGoal[i] /= 2
		}
		solve := p2Solve(newGoal, moves, cache, pCache)
		if solve != math.MaxInt {
			rec = append(rec, 2*solve+len(combo))
		}
	}

	ret := math.MaxInt
	for i := 0; i < len(rec); i++ {
		ret = min(ret, rec[i])
	}
	(*cache)[key] = ret
	return ret
}

func p2Goal2State(goal *State) State {
	ret := make([]int, len(*goal))
	for i, digit := range *goal {
		ret[i] = digit % 2
	}
	return ret
}

func patternsFor(goal State, moves []Move, pCache *map[uint32][][]Move) [][]Move {
	key := stateHash(goal) ^ movesHash(moves)
	if ret, ok := (*pCache)[key]; ok {
		return ret
	}

	var ret [][]Move
	for r := 0; r <= len(moves); r++ {
		for _, combo := range combinations(moves, r) {
			state := make(State, len(goal))
			for _, move := range combo {
				p1Transform(&state, move)
			}
			if stateEq(state, goal) {
				ret = append(ret, combo)
			}
		}
	}
	(*pCache)[key] = ret
	return ret
}

func p2Transform(state *State, move Move) {
	for _, pos := range move {
		(*state)[pos]--
	}
}

func p1Transform(state *State, move Move) {
	for _, pos := range move {
		(*state)[pos] ^= 1
	}
}

func intsPlease(s string) []int {
	return nonEmpties(s, func(s string) int { res, _ := strconv.Atoi(s); return res })
}

func nonEmpties[T any](s string, f func(string) T) []T {
	var ret []T
	for _, el := range strings.Split(s, ",") {
		if el != "" {
			ret = append(ret, f(el))
		}
	}
	return ret
}

func stateEq(s1, s2 State) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, d := range s1 {
		if d != s2[i] {
			return false
		}
	}
	return true
}

func movesHash(moves []Move) uint32 {
	k := uint32(0)
	for _, m := range moves {
		k ^= stateHash(m)
	}
	return k
}

func stateHash(s State) uint32 {
	hasher := fnv.New32()
	for _, el := range s {
		buf := make([]byte, 5)
		buf[4] = '|'
		binary.BigEndian.PutUint32(buf, uint32(el))
		_, _ = hasher.Write(buf)
	}
	return hasher.Sum32()
}

func combinations[C any](it []C, r int) [][]C {
	n := len(it)
	if r > n {
		return [][]C{}
	}

	var indices []int
	for x := range r {
		indices = append(indices, x)
	}

	take := func(xs []int) []C {
		row := make([]C, r)
		for i, x := range xs {
			row[i] = it[x]
		}
		return row
	}

	var ret [][]C
	ret = append(ret, take(indices))
	for {
		found := -1
		for i := r - 1; i >= 0; i-- {
			if indices[i] != i+n-r {
				found = i
				break
			}
		}
		if found < 0 {
			return ret
		}
		indices[found]++
		for j := found + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}
		ret = append(ret, take(indices))
	}
}
