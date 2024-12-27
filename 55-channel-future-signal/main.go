package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	sig1 := make(chan struct{})
	ch2 := make(chan string)
	sig2 := make(chan struct{})
	go sender(ch1, "sender-1", time.Millisecond*100, 5)
	go sender(ch2, "sender-2", time.Millisecond*101, 5)

	go receiver(ch1, sig1)
	go receiver(ch2, sig2)
	<-sig1
	<-sig2
}

func sender(chSend chan<- string, name string, dr time.Duration, max uint) {
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

func receiver(chRec <-chan string, sig chan<- struct{}) {
	defer func() {
		sig <- struct{}{}
	}()
	for v := range chRec {
		fmt.Println(v)
	}

}
