package main

import "fmt"

func main() {
	r1 := add(10, 20)
	r2 := add(10.5, 20.5)
	r3 := add(uint8(10), uint8(12))
	r4 := add(uint64(12312321), uint64(123432432))
	r5 := add(12312, 12312)
	fmt.Println(r1, r2, r3, r4, r5)

	var anything1 anything

	anything1 = "hello Wrold"
	str1, ok := anything1.(string)
	if ok {
		fmt.Println(str1)
	}

	anything1 = 101
	anything1 = true

	fmt.Println(anything1)
	fmt.Println(Sunday)

}

// func add(a, b any) any {

// }

// 1.18

// static dispatch , monomorphization
func add[T IType](a, b T) T {
	return a + b
}

type IType interface {
	int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

type anything interface{}

const (
	//iota   = 1
	Sunday = iota * 4
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	Sunday1    = 0
	Monday1    = 1
	Tuesday1   = 2
	Wednesday1 = 3
	Thursday1  = 4
	Friday1    = 5
	Saturday1  = 6
)
