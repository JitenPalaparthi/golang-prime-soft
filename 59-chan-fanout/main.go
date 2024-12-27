package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	for i := 1; i <= 3; i++ {
		go worker(fmt.Sprint(i), ch)
	}
	//go func() {
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- i * i
	}
	close(ch)
	//}()
	// --?
}

func worker(name string, ch <-chan int) {
	for v := range ch {
		fmt.Println("receive by worker-", name, "value-->", v)
	}
}
