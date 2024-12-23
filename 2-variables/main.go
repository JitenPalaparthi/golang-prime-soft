package main

import "fmt"

func main() {
	var num1 int // automatically zero value is given

	var num2 int = 1012312

	var num3 uint8 = 127

	var float1 float32 = 123.123

	var float2 float64 = 12312.1231

	var ok1 bool = true
	var str1 string = "Hello Prime Soft"

	var (
		ok2    bool    = false
		str2   string  = "hello minds"
		float3 float64 = 123123.12312
	)
	fmt.Println(num1, num2, num3, ok1, ok2, str1, str2, float1, float2, float3)

	var c1 complex128 = 12.12 + 1.1i
	var c2 complex64 = 12.12 + 1.1i
	fmt.Println(c1, c2)

	// expressions

	// arthemetic + -
	// logical && || !=  & |
	// comparision >= <= > < ==

	var num4 = 123 // variable creation without the data type, type inference

	var age1 = 39 // 120
	var float4 = 123.34

	str3 := "Hello PrimeSoft Minds"

	age2 := 39 // short hand declaration

	ok3 := true // bool 1 byte

	var age3 uint8 = 39

	num1 = 123 + (23 * num4 / 12) + 353 - age1*2
	ok1 = (123+(23*num4/12)+353-age1*2) > 10 && true

	c3 := complex(12.2, 1.3)

	var a, b float32 = 123.1, 1.2

	c4 := complex(a, b)

	fmt.Println(float4, str3, num4, c1, c2, c3, c4, ok3, age1, age2, age3)

	var char1 rune = 'A' + 1

	var char2 int32 = '你' + 1

	var char3 uint8 = 'A' + 1

	var char4 uint64 = 'A' + 1

	var char5 int32 = '你'

	fmt.Println(char1)
	fmt.Println(char2)
	fmt.Println(char3)
	fmt.Println(char4)
	fmt.Println(char5)

}

// numbers -> zero value is 0
// int, uint, uint8,uint16,uint32,uint64,
// int8,int16,int32,int64,float32,float64,rune,byte,uintptr
// boolean --> bool --> false
// strings --> string -> ""
// complex -> complex64, complex128 -> 0

// every variable with a type has a zero value
