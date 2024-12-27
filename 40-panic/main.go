package main

import "fmt"

func main() {
	arr := [2]int{10, 20}

	for i := 0; i <= len(arr); i++ {
		println(arr[i])
	}

	// nil deference panic
	var ptr *int

	*ptr = 123

	// division by zero panic
	divide(100)

	fmt.Println("Hello Wrold")
}

func divide(num int) {
	for i := 10; i >= 0; i-- {
		println(num / i)
	}
}
