package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	work := NewWork(wg, mu, 0)

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go work.Incr()
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go work.Decr()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Work Counter:", work.GetCounter())
}

type Work struct {
	counter int
	wg      *sync.WaitGroup
	mu      *sync.Mutex
}

func NewWork(wg *sync.WaitGroup, mu *sync.Mutex, i int) *Work {
	return &Work{wg: wg, mu: mu, counter: i}
}

func (w *Work) Incr() {
	defer w.wg.Done()
	w.mu.Lock()
	w.counter++
	w.mu.Unlock()
}

func (w *Work) Decr() {
	defer w.wg.Done()
	w.mu.Lock()
	defer w.mu.Unlock()
	w.counter--
	// if there is other logic? the whole function scope is locked since unlock is called as defer.
}

func (w *Work) GetCounter() int {
	return w.counter
}
