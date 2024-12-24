package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//var any1 any
	var str1 string = "Hello"
	//var b byte = str1[0]
	fmt.Printf("%d %d %d %d\n", &str1, &[]byte(str1)[0], &[]byte(str1)[1], &[]byte(str1)[2])

	ptr := unsafe.Pointer(&str1)
	dataPtr := *(*unsafe.Pointer)(ptr)
	fmt.Println(dataPtr)
	//ptr = ptr + 8

	v := *(*byte)(dataPtr)
	fmt.Println(string(v))

	ptr1 := uintptr(unsafe.Pointer(dataPtr))

	ptr1 = ptr1 + 1

	v = *(*byte)(unsafe.Pointer(ptr1))
	fmt.Println(string(v))
	// v := (*int)(unsafe.Pointer(ptr))
	// fmt.Printf("%x", *v)

	var num int = 100

	ptr2 := unsafe.Pointer(&num)
	fmt.Println(*(*byte)(ptr2))
	//	k := *(*unsafe.Pointer)(&num)

}
