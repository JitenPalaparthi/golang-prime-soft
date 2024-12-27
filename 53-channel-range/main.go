package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	wg := new(sync.WaitGroup)
	defer wg.Wait()
	wg.Add(4)
	go sender(wg, ch1, time.Millisecond*100, 5)
	go sender(wg, ch2, time.Millisecond*500, 5)
	go receiver(wg, ch1)
	go receiver(wg, ch2)
	// here make sure that the sender and receiver send and receive same number of values
}

func sender(wg *sync.WaitGroup, chSend chan int, dr time.Duration, max uint) {
	defer wg.Done()
	var i uint
	for {
		time.Sleep(dr)
		chSend <- int(i * i)
		if i == max {
			close(chSend)
			return
		}
		i++
	}
}

func receiver(wg *sync.WaitGroup, chRecive chan int) {
	defer wg.Done()
	for v := range chRecive { // range loop automatically exit when the channel is closed
		println(v)
	}
	fmt.Println("channel is closed")
}
