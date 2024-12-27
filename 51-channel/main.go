package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int)
	wg := new(sync.WaitGroup)
	defer wg.Wait()
	wg.Add(3)
	go sender(wg, ch, time.Millisecond*100, 1, 10)
	go sender(wg, ch, time.Millisecond*100, 11, 20)
	go receiver(wg, ch, 1, 20)
	// here make sure that the sender and receiver send and receive same number of values
}

func sender(wg *sync.WaitGroup, chSend chan int, dr time.Duration, i1, i2 int) {
	for i := min(i1, i2); i <= max(i1, i2); i++ {
		time.Sleep(dr)
		chSend <- i * i
	}
	wg.Done()
}

func receiver(wg *sync.WaitGroup, chRecive chan int, i1, i2 int) {
	for i := min(i1, i2); i <= max(i1, i2); i++ {
		fmt.Println(<-chRecive)
	}
	wg.Done()
}
