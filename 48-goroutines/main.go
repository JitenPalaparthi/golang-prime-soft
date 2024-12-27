package main

import (
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(4)
	go EvenNumbers(wg, 1, 10)

	//wg.Add(1)
	go EvenNumbers(wg, 1000, 1010)

	//wg.Add(1)
	go func(wg *sync.WaitGroup) {
		OddNumbers(2000, 2010)
		wg.Done()
	}(wg)

	//wg.Add(1)
	EvenNumbers(wg, -30, -40)
	//go OddNumbers(-10, -3000)
	// the state or counter is 3
	wg.Wait() // waiting for, the counter of wg to become zero
}

func EvenNumbers(wg *sync.WaitGroup, f, l int) {
	mi1 := min(f, l)
	mx1 := max(f, l)
	for i := mi1; i <= mx1; i++ {
		if i%2 == 0 {
			print(i, " ")
		}
	}

	wg.Done()
}

func OddNumbers(f, l int) {
	mi1 := min(f, l)
	mx1 := max(f, l)
	for i := mi1; i <= mx1; i++ {
		if i%2 != 0 {
			print(i, " ")
		}
	}
}
