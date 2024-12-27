package main

import "time"

func main() {
	ch := Generator(50, time.Millisecond*100)
	// sig := Receive(ch)
	// <-sig
	<-Receive(ch)
}

func Generator(num int, d time.Duration) chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= num; i++ {
			time.Sleep(d)
			ch <- i * i
		}
		close(ch)
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
