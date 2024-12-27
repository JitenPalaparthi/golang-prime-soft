package main

import "fmt"

func main() {
	defer println("end of main-1")
	defer println("end of main-2")

	defer func() {
		a, b, c := 10, 20, 0
		c = a + b
		println(c)
	}()

	a := 100
	b := 100

	defer func(a int) {
		println("a in defer:", a)
	}(a)

	defer func() {
		println("b in defer:", b)
	}()

	a++
	b++

	println("a normal exec:", a)
	println("b normal exec:", b)

	str := "Hello PrimeSoft"

	defer println()
	for _, c := range str {
		defer print(string(c))
	}

	fmt.Println()

	println("start of main")
}

// defer , defers the execution to the end of the call stack of a caller
// defer can only be called on funcs or methods
