package main

import (
	"runtime"
	"time"
)

func main() {
	ch := Generator(time.Second*3, time.Millisecond*100)
	// sig := Receive(ch)
	// <-sig
	<-Receive(ch)
}

func Generator(duration time.Duration, sleep time.Duration) chan int {
	ch := make(chan int)
	//sig := Timeout(duration)
	sig := time.After(duration)
	go func() {
		i := 1
		//out:
		for {
			select {
			case ch <- i * i:
				time.Sleep(sleep)

			case <-sig:
				close(ch)
				runtime.Goexit()
			}
			i++
		}
	}()
	return ch
}

func Receive(ch chan int) chan struct{} {
	sig := make(chan struct{})
	go func() {
		for v := range ch {
			println(v)
		}
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}

func Timeout(d time.Duration) chan struct{} {
	sig := make(chan struct{})
	go func() {
		time.Sleep(d)
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}
