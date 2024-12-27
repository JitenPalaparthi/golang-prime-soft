package main

import (
	"fmt"
	"os"
)

func main() {
	defer println("End of main")
	num := 100
	func() { // func1
		defer println("end of func1")
		defer func() {
			if r := recover(); r != nil {
				// handle it here
				fmt.Println(r, "recover from func-1 panic")
			}
		}()
		func() { // func2
			defer recoverme()
			defer println("end of func-2")
			println("start of func-2")
			for i := 10; i >= 0; i-- {
				print(num/i, " ") // panics
			}
			defer println("defer after panic --???")
			println("graceful exit-->?")
		}()
		//Fatal("just exit with non zero") , fatal never calles the deferred calls..
		panic("just panic to see what would happen")
		println("do something in func1")
	}()

	str := "hello Wrold"
	for _, v := range str {
		println(string(v))
	}
	println("graceful exit of main-->?")
}

func Fatal(msgs ...any) {
	fmt.Println(msgs...)
	os.Exit(1) // exiting the application abruptly is fatal.
	// once os.Exit is called , cannot recover from it
}
func recoverme() {
	if r := recover(); r != nil {
		// handle it here
		fmt.Println("recover from the panic.So cascading crash should not happen.-->", r)
	}
}
