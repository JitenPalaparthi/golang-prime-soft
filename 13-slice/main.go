package main

import "fmt"

func main() {

	//slice1 := []int{1, 2, 3, 4, 5} // cap 5 len 5
	//PrintSliceHeader(slice1, "slice1")
	slice1 := make([]int, 5, 10) // cap 10 len 5
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 1, 2, 3, 4, 5
	//PrintSliceHeader(slice1, "slice1")
	fmt.Println(slice1)
	num := 100
	double1(num)
	println(num)

	double2(slice1)
	fmt.Println("After double")
	fmt.Println(slice1)
	slice2 := slice1 // header copy
	slice2[0] = 99999
	slice2[1] = 88888
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	slice2 = append(slice2, 111111)
	slice2[2] = 7777
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
}

func double1(n int) {
	n = n * 2
}

func double2(slice []int) {
	for i, v := range slice {
		slice[i] = v * 2
	}
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
