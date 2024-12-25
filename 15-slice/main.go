package main

import "fmt"

func main() {

	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	slice2 := slice1[:] // shallow copy

	slice3 := slice1[:5]
	slice4 := slice1[2:5]
	slice5 := slice1[5:]

	PrintSliceHeader(slice1, "slice1")
	PrintSliceHeader(slice2, "slice2")
	PrintSliceHeader(slice3, "slice3")
	PrintSliceHeader(slice4, "slice4")
	PrintSliceHeader(slice5, "slice5")

	slice6 := append(slice1[:2], slice1[3:]...)
	fmt.Println(slice6)

	// make, append , cap, len, copy, clear

	slice7 := make([]int, 20)
	//var slice7 []int
	copy(slice7, slice1[:5]) // deep copy
	fmt.Println(slice7)
	clear(slice7)
	fmt.Println(slice7)
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

func deleteElem(slice []int, index int) ([]int, error) {

	return nil, nil
}
