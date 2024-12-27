package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 100
		close(ch)
	}()

	sig := make(chan struct{})
	go func() {
		for i := 1; i <= 5; i++ {
			select {
			case v, ok := <-ch:
				println(v, ok)
			}
		}
		sig <- struct{}{}
	}()
	<-sig

	server1 := make(chan string)
	server2 := make(chan string)
	server3 := make(chan string)

	go getServer(server1, "server-1", time.Millisecond*100)
	go getServer(server2, "server-2", time.Millisecond*100)
	go getServer(server3, "server-3", time.Millisecond*100)

	select {
	case <-server1:
		fmt.Println("server1 is going to be used")
	case <-server2:
		fmt.Println("server2 is going to be used")
	case <-server3:
		fmt.Println("server3 is going to be used")
	}

}

func getServer(ch chan string, name string, d time.Duration) {
	time.Sleep(d)
	ch <- name
}
