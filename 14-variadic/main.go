package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println("hello", "World")
	fmt.Println("hello", "World", "How", "Are", "You")

	slice1 := make([]int, 2)
	fmt.Println(slice1, "cap and len", cap(slice1), len(slice1))
	slice1[0], slice1[1] = 100, 101 // len 2 cap 2

	slice1 = append(slice1, 102, 103, 104, 105, 106)
	// slice1 = append(slice1, 102)
	// fmt.Println(slice1, "cap and len", cap(slice1), len(slice1))
	// slice1 = append(slice1, 103)
	// fmt.Println(slice1, "cap and len", cap(slice1), len(slice1))
	// slice1 = append(slice1, 104)
	// fmt.Println(slice1, "cap and len", cap(slice1), len(slice1))
	// slice1 = append(slice1, 105)
	// fmt.Println(slice1, "cap and len", cap(slice1), len(slice1))
	// slice1 = append(slice1, 106)
	fmt.Println(slice1, "cap and len", cap(slice1), len(slice1))

	// var nums ...int

	fmt.Println(SumOf())
	fmt.Println(SumOf(10))
	fmt.Println(SumOf(10, 20))
	fmt.Println(SumOf(10, 20, 123, 43, 34, 43, 6, 56, 7, 8))
	println(SumOf(slice1...)) // a slice can be passed as an argument for variadic like this
	println(SumOfElenms(slice1))
	println(SumOfElenms([]int{10}))
	println(SumOfElenms([]int{10, 20}))
	slice2 := []int{10, 20, 30}
	println(SumOfElenms(slice2))

	arr := [3]int{10, 11, 12} // an array can be converted to slice
	println(SumOf(arr[:]...)) // a slice can be passed as an argument for variadic like this

	slice3 := arr[:] // array is convered to slice
	println(SumOfElenms(arr[:]))
	fmt.Println("arr:", arr)
	slice3[0] = 999999
	fmt.Println("arr:", arr)
	slice3 = append(slice3, 1111)
	slice3[1] = 8888
	fmt.Println("arr:", arr)

	slice4 := []int{1, 2, 3}
	fmt.Println(slice4)
	slice4 = AddAndDouble(slice4, 4, 5, 6)
	fmt.Println(slice4)
}

func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func SumOfElenms(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func AddAndDouble(slice []int, num ...int) []int {
	slice = append(slice, num...)
	for i, v := range slice {
		slice[i] = v * 2
	}
	return slice
}

// ...type variadic parameter
// the variadic parameter must be the last parameter
// can call only one variadic parameter for a function or method
// variadic parameter can only be used in funcs and methods
// can use range loop for variadic param
