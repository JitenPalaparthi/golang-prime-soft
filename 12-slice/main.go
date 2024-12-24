package main

import (
	"fmt"
	v2 "math/rand/v2"
)

func main() {

	var slice1 []int // slice declaration

	if slice1 == nil {
		println("slice1 is nil")
	}
	PrintSliceHeader(slice1, "slice1")

	slice1 = make([]int, 5) // instantiating a slice
	PrintSliceHeader(slice1, "slice1")

	slice2 := []int{}
	PrintSliceHeader(slice2, "slice2")

	// slice2 := make([]int, 5, 10) // short hand declarion
	// PrintSliceHeader(slice2, "slice2")

	// slice3 := []int{1, 2, 3, 4, 5} // cap :5 len :5
	// PrintSliceHeader(slice3, "slice3")

	//for i, _ := range slice1 {
	for i := range slice1 { // range iterate thru the length
		slice1[i] = v2.IntN(100)
	}
	PrintSliceHeader(slice1, "slice1")
	// length ? 5 cap: 5

	slice1 = append(slice1, 9999)
	PrintSliceHeader(slice1, "slice1")
	slice1 = append(slice1, 1111)

}

func PrintSliceHeader(slice []int, name string) {
	fmt.Println("----------------------------")
	fmt.Println(name)
	fmt.Println("Elements->", slice)
	fmt.Printf("Address of the slice Header->%p\n", &slice)
	if len(slice) > 0 {
		fmt.Printf("Pointer of the slice->%p\n", &slice[0])
	}
	fmt.Printf("len:%d cap: %d \n ", len(slice), cap(slice))
	//fmt.Println("----------------------------")
}

// slice
// slice is collection of elements of same type
// the slice can be grown at runtime
// slice has a header
// make (builtin func) is used to instantiate a slice
// slice has len and cap
// append is a built in function that is only for slice
