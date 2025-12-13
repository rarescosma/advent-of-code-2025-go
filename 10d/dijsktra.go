package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

type P1State = int

type Move []int

type Item[S comparable] struct {
	state S
	cost  int
	index int
}

// PriorityQueue implementation copied verbatim from
// https://pkg.go.dev/container/heap#example-package-PriorityQueue
type PriorityQueue[S comparable] []*Item[S]

func (pq PriorityQueue[S]) Len() int { return len(pq) }

func (pq PriorityQueue[S]) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue[S]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[S]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[S])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[S]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func main() {
	file, _ := os.Open("inputs/10.in")

	scanner := bufio.NewScanner(file)
	p1 := 0
	for scanner.Scan() {
		line := scanner.Text()[1:]
		line = line[:len(line)-1]
		parts := strings.Split(line, "] (")
		p1Goal := 0
		for i, c := range parts[0] {
			if c == '#' {
				p1Goal = p1Goal + (1 << i)
			}
		}

		parts = strings.Split(parts[1], ") {")

		var p1moves []Move
		for _, part := range strings.Split(parts[0], ") (") {
			p1moves = append(p1moves, intsPlease(part))
		}

		p1 += dijsktra(0, p1Goal, p1moves, transformP1)
	}

	println("p1:", p1)
}

func transformP1(state *P1State, move Move) P1State {
	ret := *state
	for _, pos := range move {
		ret ^= 1 << pos
	}
	return ret
}

func dijsktra[S comparable](
	state S,
	goal S,
	moves []Move,
	transform func(s *S, m Move) S,
) int {
	known, pq := make(map[S]int), make(PriorityQueue[S], 1)

	pq[0] = &Item[S]{state: state, cost: 0, index: 0}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item[S])

		if item.state == goal {
			return item.cost
		}

		for _, move := range moves {
			newCost := item.cost + 1
			newState := transform(&item.state, move)

			if oldCost, ok := known[newState]; (ok && newCost < oldCost) || !ok {
				known[newState] = newCost
				item := &Item[S]{state: newState, cost: newCost}
				heap.Push(&pq, item)
			}
		}
	}
	return -1
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
