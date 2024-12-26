package main

import "fmt"

func main() {

	funcMap := make(map[string]any)

	//a, b := 10, 20

	funcMap["add"] = func(a, b int) int {
		return a + b
	}

	subFn := func(a, b int) int {
		return a - b
	}

	funcMap["sub"] = subFn
	funcMap["mul"] = mul

	var funcDiv func(int, int) int = func(i1, i2 int) int {
		return i1 / i2
	}

	funcMap["div"] = funcDiv

	funcMap["greet"] = func() {
		fmt.Println("Hello PrimeSoft")
	}

	funcMap["sq"] = func(a int) int {
		return a * a
	}

	funcMap["empty"] = struct{}{}
	funcMap["num"] = 12321
	funcMap["rect"] = Rect{L: 123.23, B: 123.23}
	funcMap["intSLice"] = []int{10, 20, 30}

	// 1. func(int,int)int
	// 2. func(int)int
	// 3. func()

	a, b, c := 10, 20, 30

	for key, val := range funcMap {
		switch fn := val.(type) {
		case func():
			fmt.Print(key, "= executing the following function.--> ")
			fn()
		case func(int) int:
			r := fn(c)
			fmt.Println(key, "of ", c, "=", r)
		case func(int, int) int:
			r := fn(a, b)
			fmt.Println(key, "of ", a, b, "=", r)
		case struct{}:
			fmt.Println(key, "is empty struct")
		case int:
			fmt.Println(key, "is type int. Value is", fn)
		case Rect:
			fmt.Println(key, "values are-->", fn)
		case []int:
			fmt.Println(key, "values are", fn)
		default:
			fmt.Println(key, "of", fn, "is not implemented")
		}

	}

}

type Rect struct {
	L, B float32
}

func mul(a, b int) int {
	return a * b
}

// 1. calling functions with type switch
// 2. closures
