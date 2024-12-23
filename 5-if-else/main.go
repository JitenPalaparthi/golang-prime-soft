package main

import (
	"fmt"
	"strconv"
)

var global int = 100 // data segment

const PI float32 = 3.14 //code segment

func main() {
	var age uint8 = 17 // stack memory

	const (
		MAX    = 99
		MIN    = MAX * 10
		SECOND = 60
		MINUTE = 60 * SECOND
		HOUR   = 24 * MINUTE * SECOND
	)

	//var any1 any = "hello World" // heap memory

	{
		b := 200
		c := b * 2
		println(b, c)
	}

	fmt.Println(age >= 18 && (true || false) && true)

	//if age >= 18 && (true || false) && true {
	if age >= 18 {
		// false && true && true
		// false && true
		// false
		c := 100
		println("eligible to vote", c)
	} else {
		println("not eligible to vote")
	}

	//println(c)

	age = 18

	g := 'f' // it is rune which is nothing but int32
	if age >= 18 && (g == 102 || g == 'F') {
		println("she is elibible for marriage becasue age is", age)
	} else if age >= 21 && (g == 'M' || g == 'm') {
		println("he is elibible for marriage becasue age is", age)
	} else {
		println("not eligible for marraige")
	}

	str1 := "12312"

	if num, err := strconv.Atoi(str1); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num)
	}
	//println(num)
}
