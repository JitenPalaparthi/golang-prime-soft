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
	wg.Add(2)
	go sender(wg, ch, time.Millisecond*100, 20)
	go receiver(wg, ch)
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
	for {
		if v, ok := <-chRecive; ok {
			fmt.Println(v)
		} else {
			fmt.Println("channel is clsed so exiting")
			break
		}
	}

}
