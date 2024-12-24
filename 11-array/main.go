package main

import (
	"fmt"
	v2 "math/rand/v2"
	"reflect"
)

func main() {
	// var num1 int // ?
	// const PI float32 = 3.14
	var arr1 [5]int  //? what is the zero value of arr1? [0 0 0 0 0]
	var arr2 [3]int  //? [0 0 0]
	var arr3 [2]bool //? [false false]

	fmt.Println("Type of arr1:", reflect.TypeOf(arr1), arr1)
	fmt.Println("Type of arr2:", reflect.TypeOf(arr2), arr2)
	fmt.Println("Type of arr3:", reflect.TypeOf(arr3), arr3)

	arr1[0] = 10
	arr1[1] = 20

	for i := 0; i < len(arr1); i++ {
		arr1[i] = v2.IntN(100)
	}

	fmt.Println("Type of arr1:", reflect.TypeOf(arr1), arr1)

	arr4 := [3]int{10, 11, 12}   // short hand declaration of array
	arr5 := [...]int{10, 11, 12} // short hand declaration of array. The length is evaluated by the compiler

	var arr6 [3]int = arr5 // both arr5 and arr6 are same type
	// this is deep copy of an array

	//var arr7 [3]int = ([3]int)(arr1) // two different types

	fmt.Println("Type of arr4:", reflect.TypeOf(arr4), arr4)
	fmt.Println("Type of arr5:", reflect.TypeOf(arr5), arr5)
	fmt.Println("Type of arr6:", reflect.TypeOf(arr6), arr6)

	// arr1

	for _, v := range arr1 {
		println(v)
	}

	// copy from one array to another

	var arr8 [3]int
	for i := 0; i < min(len(arr1), len(arr8)); i++ {
		arr8[i] = arr1[i]
	}

	fmt.Println("Type of arr8:", reflect.TypeOf(arr8), arr8)

	s1 := SumOf(arr1)
	fmt.Println("Sum fo arr1:", s1)
	SumOf3(arr8)

	arr2d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	/*
		[
			1 2 3
			4 5 6

			1 4
			2 5
			3 6
		]
	*/

	for _, a1 := range arr2d {
		for _, v1 := range a1 {
			fmt.Print(v1, " ")
		}
		println()
	}

	arr3d := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}

	fmt.Println("3d array")
	for _, a1 := range arr3d {
		for _, a2 := range a1 {
			for _, v1 := range a2 {
				fmt.Print(v1, " ")
			}
			println()
		}
	}
}

// what is an array
// the length of an array is fixed, should be known to the compiler
// cannpot type cast arrays directly

func SumOf(arr [5]int) int {
	sum := 0

	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumOf3(arr [3]int) int {
	sum := 0

	for _, v := range arr {
		sum += v
	}
	return sum
}
