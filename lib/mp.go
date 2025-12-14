package lib

import (
	"math"
	"sync"
)

type Mp[I, A any] struct {
	numThreads int
	items      []I
	worker     func(items []I) A

	_wg        *sync.WaitGroup
	_chunkSize int
	_ansChan   chan A
}

func MakeMp[I, A any](numThreads int, items []I, worker func(chunk []I) A) Mp[I, A] {
	var _wg sync.WaitGroup
	_wg.Add(numThreads)
	_chunkSize := int(math.Ceil(float64(len(items)) / float64(numThreads)))
	_ansChan := make(chan A)

	return Mp[I, A]{numThreads, items, worker, &_wg, _chunkSize, _ansChan}
}

func (mp *Mp[I, A]) Go() chan A {
	go func() {
		mp._wg.Wait()
		close(mp._ansChan)
	}()

	for n := range mp.numThreads {
		start := min(n*mp._chunkSize, len(mp.items))
		end := min((n+1)*mp._chunkSize, len(mp.items))
		chunk := mp.items[start:end]

		if len(chunk) == 0 {
			mp._wg.Done()
			continue
		}

		go func(chunk []I, wg *sync.WaitGroup) {
			defer wg.Done()
			mp._ansChan <- mp.worker(chunk)

		}(chunk, mp._wg)
	}
	return mp._ansChan
}
