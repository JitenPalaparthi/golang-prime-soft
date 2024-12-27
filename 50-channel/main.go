package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

func main() {
	var ch chan int
	if ch == nil {
		fmt.Println("nil channel")
		fmt.Println("Size of chan-->", unsafe.Sizeof(ch))
	}
	ch = make(chan int) // unbuffered channel
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		time.Sleep(time.Second * 1)
		ch <- 100
		wg.Done()
	}()
	go func() {
		v := <-ch // how long it is blocked?
		println(v)
		wg.Done()
	}()
	wg.Wait()
}

// channel --> conduit -- kind of a pipe or a tube or a queue
// do not communicate by sharing memory;
//  share memory by communicating
// some buffer, send queue, receive queue and mutex
// how to create a channel?
// the sender is blocked untile the receiver receives the value
// the receiver is blocked until the seder sends the value
// to send a value to the channel ch <- 100
// to receive a value v:= <- ch
// the order of sender and receiver of goroutines does not matter, as long as they comply with unblocking rules
