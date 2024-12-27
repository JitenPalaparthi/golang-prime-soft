package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch1 := Generator("sender-1", time.Second*3, time.Millisecond*100)
	ch2 := Generator("sender-2", time.Second*3, time.Millisecond*100)
	Receiver(ch1, ch2)
	time.Sleep(time.Second * 10) // insted of sleep use a better mech to gracefully deal with the main exit
}

func Generator(name string, duration time.Duration, sleep time.Duration) chan string {
	ch := make(chan string)
	//sig := Timeout(duration)
	sig := time.After(duration)
	go func() {
		i := 1
		//out:
		for {
			select {
			case ch <- fmt.Sprint(name, "--->", i*i):
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

func Receiver(chs ...chan string) {
	for _, ch := range chs {
		go func(ch chan string) {
			for v := range ch {
				println(v)
			}

		}(ch)
	}
}
