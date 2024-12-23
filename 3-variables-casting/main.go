package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func main() {

	var num1 uint8 = 100
	var num2 int = int(num1)
	println(num2)

	//ok1 := true

	//var num3 uint8 = ok1 // no possible bcz bool cannot be casted to number

	var num3 uint64 = 12312
	var num4 uint8 = uint8(num3)
	var num5 uint16 = uint16(num3)
	fmt.Printf("num3 : %b-->%d\n", num3, num3)
	fmt.Printf("num4 : %b-->%d\n", num4, num4)
	fmt.Printf("num5 : %b-->%d\n", num5, num5)

	var float1 float64 = 123123123123.123123

	var num6 int = int(float1)
	fmt.Printf("num6 : %b-->%d\n", num6, num6)

	var char1 rune = 'A' + 1

	var char2 int32 = '你' + 1

	var char3 uint8 = 'A' + 1

	var char4 uint64 = 'A' + 1

	var char5 int32 = '你'

	var char6 float32 = 'A'

	var float2 float64 = 123
	fmt.Println(string(char1))
	fmt.Println(string(char2))
	fmt.Println(string(char3))
	fmt.Println(string(char4))
	fmt.Println(string(char5))
	fmt.Println(char6)

	fmt.Println(float2)

	fmt.Println("Value:", float2, "Size of", unsafe.Sizeof(float2), "Type:", reflect.TypeOf(float2))

	a1, p1 := AreaAndPerimeter(123.123, 123.32)
	fmt.Println("Area:", a1, "Perimeter:", p1)

	a2, _ := AreaAndPerimeter(123.123, 123.32)
	fmt.Println("Area:", a2)

	_, p2 := AreaAndPerimeter(123.123, 123.32)
	fmt.Println("Perimter:", p2)

	var str1 = "112312"

	// type casting vs conversion
	var num7 int

	num7, _ = strconv.Atoi(str1)
	fmt.Println(num7)

	var str2 = "1124er312"

	// type casting vs conversion
	var num8 int
	//var err error
	num8, err := strconv.Atoi(str2)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(num8)
	}

	str3 := "123123.123123"

	float3, err := strconv.ParseFloat(str3, 64)

	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(float3)

		fmt.Printf("%.4f\n", float3)
	}

}

func AreaAndPerimeter(l float32, b float32) (a float32, p float32) {
	return l * b, 2 * (l + b)
}

// any number type can be casted to another number type
//
