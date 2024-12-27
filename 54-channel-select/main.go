package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	wg := new(sync.WaitGroup)
	defer wg.Wait()
	// wg.Add(3)
	// go sender(wg, ch1, "sender-1", time.Millisecond*100, 5)
	// go sender(wg, ch2, "sender-2", time.Millisecond*500, 5)
	// go receiver(wg, ch2, ch1)

	wg.Add(3)
	go sender(wg, ch1, "sender-1", time.Millisecond*100, 5)
	go sender(wg, ch2, "sender-2", time.Millisecond*500, 5)
	go receiverR(wg, ch2, ch1) // here make sure that the sender and receiver send and receive same number of values
}

func sender(wg *sync.WaitGroup, chSend chan<- string, name string, dr time.Duration, max uint) {
	defer wg.Done()
	var i uint
	for {
		time.Sleep(dr)
		chSend <- fmt.Sprint(name, "--->", i*i)
		if i == max {
			close(chSend)
			return
		}
		i++
	}
}

func receiver(wg *sync.WaitGroup, ch1, ch2 <-chan string) {
	defer wg.Done()
	done1, done2 := false, false
	for {
		// v1 := <-ch1
		// println(v1)
		// v2 := <-ch2
		// println(v2)
		if done1 && done2 {
			fmt.Println("finished reading from both the channels")
			return
		}
		select { // select works only with channels
		case v, ok := <-ch1:
			if ok {
				fmt.Println(v)
			} else {
				done1 = true
			}
		case v, ok := <-ch2:
			if ok {
				fmt.Println(v)
			} else {
				done2 = true
			}
		}
	}
}

func receiverR(wg *sync.WaitGroup, ch1, ch2 <-chan string) {
	wg.Add(1)
	go func() {
		for v := range ch1 {
			fmt.Println(v)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for v := range ch2 {
			fmt.Println(v)
		}
		wg.Done()
	}()

	wg.Done()
}
