package lib

import (
	"math"
	"sync"
)

type Pool[I, A any] struct {
	numThreads int
	items      []I
	worker     func(items []I) A

	_wg        *sync.WaitGroup
	_chunkSize int
	_ansChan   chan A
}

func NewPool[I, A any](numThreads int, items []I, worker func(chunk []I) A) Pool[I, A] {
	var _wg sync.WaitGroup
	_wg.Add(numThreads)
	_chunkSize := int(math.Ceil(float64(len(items)) / float64(numThreads)))
	_ansChan := make(chan A)

	return Pool[I, A]{numThreads, items, worker, &_wg, _chunkSize, _ansChan}
}

func (mp *Pool[I, A]) Go() chan A {
	go func() {
		mp._wg.Wait()
		close(mp._ansChan)
	}()

	for n := range mp.numThreads {
		start := min(n*mp._chunkSize, len(mp.items))
		end := min((n+1)*mp._chunkSize, len(mp.items))

		if end <= start {
			mp._wg.Done()
			break
		}

		go func(chunk []I, wg *sync.WaitGroup) {
			defer wg.Done()
			mp._ansChan <- mp.worker(chunk)
		}(mp.items[start:end], mp._wg)
	}
	return mp._ansChan
}
