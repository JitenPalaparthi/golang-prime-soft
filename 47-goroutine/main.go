package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//fmt.Println("Initially GoRoutines", runtime.NumGoroutine())
	runtime.Goexit()
	go println("Hello Primesoft")
	go func() {
		go func() {
			i := 1
			for ; ; i++ {
				if i == 100 {
					//return
					runtime.Goexit()
				}
				if i%2 == 0 {
					//fmt.Println("Number of GoRoutines", runtime.NumGoroutine())
					go func(a int) {
						time.Sleep(time.Millisecond * 100)
						println(a)
					}(i)
				}
			}
		}()
	}()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		time.Sleep(time.Second * 2)
		panic("timed out")
	}()
	println("End of main")
	//fmt.Println("Number of GoRoutines", runtime.NumGoroutine())
	//time.Sleep(time.Second * 10)
	//runtime.Goexit()
	//os.Exit(2)

	// in main , it returns abruptly with a non zero exit
}

// What is go routine? --> Simple small green thread
// 1. main is also a go routine
// 2. by default no goroutine waits for the othere goroutine to completes its execution
// 3. developer cannot decide the order of go routine execution.Bcz concurrently
