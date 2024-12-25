package main

import "fmt"

func main() {

	num := 100
	incr(num) // call by value
	println(num)
	incrR(&num) // call by reference or address
	println(num)

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	sq(slice1, 4, 5, 6)
	sqR(&slice2, 4, 5, 6)

	fmt.Println("slice1->", slice1)
	fmt.Println("slice2->", slice2)

}

func incr(n int) {
	n++
}

func incrR(n *int) {
	if n != nil {
		*n++
	}
}

func sq(slice []int, nums ...int) {
	slice = append(slice, nums...)
	if slice != nil {
		for i, v := range slice {
			slice[i] = v * v
		}
	}
}

func sqR(slice *[]int, nums ...int) {
	if slice != nil {
		*slice = append(*slice, nums...)
		for i, v := range *slice {
			(*slice)[i] = v * v
		}
	}
}

// there is no direct pointer arthimetic in go
