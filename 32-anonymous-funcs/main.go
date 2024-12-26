package main

import "fmt"

func main() {

	var s1 struct {
		Name    string
		Message string
		Greet   func(string, string)
	} = struct {
		Name    string
		Message string
		Greet   func(string, string)
	}{
		Name:    "PrimeSoft",
		Message: "Hello",
		Greet: func(s1, s2 string) {
			println(s1, s2)
		},
	}

	s1.Greet(s1.Message, s1.Name)

	// 1.20 >=

	//funcs := make([]func(int, int) int, 0)

	fmt.Println("funcmap and funcs slice")
	funcMap := make(map[int]func(int, int) int)

	funcMap[1] = func(i1, i2 int) int {
		return i1 + i2
	}
	funcMap[2] = func(i1, i2 int) int {
		return i1 - i2
	}
	funcMap[3] = func(i1, i2 int) int {
		return i1 * i2
	}
	funcMap[4] = func(i1, i2 int) int {
		return i1 / i2
	}
	funcMap[5] = func(i1, i2 int) int {
		return i1 % i2
	}

	var funcs1 []func(int, int) int
	for i := 1; i <= 5; i++ {
		funcs1 = append(funcs1, funcMap[i])
	}

	for _, v := range funcs1 {
		r := v(10, 20)
		fmt.Println(r)
	}

	fmt.Println("loop closure global variable")

	var funcs2 []func()

	for i := 1; i <= 5; i++ {
		funcs2 = append(funcs2, func() {
			fmt.Println("Global variable is:", i)
		})
	}

	for _, v := range funcs2 {
		v()
	}

	var funcs3 []func()

	i := 1
loop:
	funcs3 = append(funcs3, func() {
		fmt.Println("Global variable is:", i)
	})
	i++
	if i < 5 {
		goto loop
	}

	for _, v := range funcs3 {
		v()
	}

}
