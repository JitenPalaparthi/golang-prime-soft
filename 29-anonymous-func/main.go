package main

import (
	"fmt"
	"unsafe"
)

func main() {

	func() {
		println("Hello Primesoft")
	}() // creating and executing the function

	f1 := func() {
		println("Hello Primesoft")
	} // creating a function and assigning it to a variable of type func

	var f2 func() // declaring a variable
	if f2 == nil {
		fmt.Println("f2 is nil", "Size of:", unsafe.Sizeof(f2))
	}

	f2 = func() {
		println("Hello Primesoft")
	} // here assigning a function to a variable which has been declared

	if f2 == nil {
		fmt.Println("f2 is nil", "Size of:", unsafe.Sizeof(f2))
	}

	f1() // exexuting the function
	f2() // executing the function

	c1 := func(a, b int) int {
		return a + b
	}(10, 20) // executor

	f3 := func(a, b int) int {
		return a + b
	}

	c2 := f3(10, 20)

	fmt.Println(c1, c2)

	f4 := func(a, b int) (int, int) {
		add := func(a1, b1 int) int {
			return a1 + b1
		}
		sub := func() int {
			return a - b
		}
		return add(a, b), sub()
	}
	a1, s1 := f4(10, 20)
	fmt.Println(a1, s1)

	a2, s2 := func(a, b int) (int, int) {
		return a + b, a - b
	}(10, 20)
	fmt.Println(a2, s2)
}
