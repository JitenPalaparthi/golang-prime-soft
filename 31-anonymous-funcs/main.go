package main

import "fmt"

var globalFn func() = func() {
	fmt.Println("Hello Primesoft")
}

func main() {

	globalFn()

	// a, b := 10, 20

	// r := func() int {
	// 	return a + b
	// }
	r1 := calc(10, 20, func(i1, i2 int) int {
		return i1 + i2
	})

	r2 := calc(10, 20, func(i1, i2 int) int {
		return i1 - i2
	})

	r3 := calc(10, 20, func(i1, i2 int) int {
		return i1 * i2
	})

	r4 := calc(10, 20, func(i1, i2 int) int {
		return i1 / i2
	})

	r5 := calc(11, 2, rem)
	fmt.Println(r1, r2, r3, r4, r5)

	var fn1 Fn = func(x int) int {
		return x
	}

	r6 := fn1(100)

	r7 := fn1.Sq(100)
	fmt.Println(r6, r7)

	f := Greet("Hello", "PrimeSoft")
	f()

}

func calc(a, b int, fn func(int, int) int) int {
	if a == 0 {
		a = 1
	}
	if b == 0 {
		b = 1
	}
	return fn(a, b)
}

func rem(a, b int) int {
	return a % b
}

type Fn func(x int) int

func (f Fn) Sq(x int) int {
	return f(x) * f(x)
}

func Greet(name, message string) func() {
	return func() {
		fmt.Println(name, message)
	}
}
