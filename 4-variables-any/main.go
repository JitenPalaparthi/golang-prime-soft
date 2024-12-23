package main

import "fmt"

func main() {

	var num1 int = 12312

	var ok = true

	str1 := "Hello World"

	var any1 any // what is the zero value of any ? nil

	if any1 == nil {
		fmt.Println(any1)
	}
	//var str1 string

	var any2 any = "hello WOrld"

	var str2 string = any2.(string) // any.(type)
	fmt.Println(str2)
	any1 = str1

	any1 = ok

	var ok2 bool = any1.(bool)
	println(ok2)
	any1 = num1

	num3 := any1.(int)
	println(num3)

	any1 = 1232.1231231
	float3, ok3 := any1.(float64)
	if ok3 {
		fmt.Println(float3)
	} else {
		println("could not assert")
	}

	any1 = any2

	//Object obj =123
}

// empty interface --> interface{} or any
