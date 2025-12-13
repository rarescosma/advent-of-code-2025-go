package lib

import "container/heap"

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

func Dijsktra[S comparable, M any](
	state S,
	goal S,
	moves []M,
	transform func(s *S, m M) S,
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
