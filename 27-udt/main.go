package main

import (
	"fmt"
	"reflect"
)

//type integer = int // alias for existing type

type myint1 int // new type

func (m myint1) ToString() string {
	return fmt.Sprint(m)
}

func (m myint1) Sq() myint1 {
	return m * m
}

type myint2 int

func (m myint2) Cube() myint2 {
	return m * m * m
}

type myint3 myint2

func main() {
	var m1 myint1 = 12312
	str1 := m1.ToString()
	sq1 := m1.Sq()
	cb1 := myint2(m1).Cube()
	fmt.Println(str1, sq1, cb1, "Type of m1:", reflect.TypeOf(m1))

	var num1 int = 123
	str2 := myint1(num1).ToString()
	sq2 := myint1(num1).Sq()
	cb2 := myint2(num1).Cube()
	fmt.Println(str2, sq2, cb2)

	var float1 float32 = 123.23

	str3 := myint1(float1).ToString()
	sq3 := myint1(float1).Sq()
	cb3 := myint2(float1).Cube()
	fmt.Println(str3, sq3, cb3)

	var m2 myint2
	fmt.Println(m2)

	// call ToString, Sq and Cube methods on m2
	var m3 myint3
	fmt.Println(m3)
	// call ToString, Sq and Cube methods on m3

}

// something about format using Sprint
//str2 := fmt.Sprint(true, "12321.123", 12312.123, 14343, "some value")

// num1 := 100
// str3 := strconv.Itoa(num1)

// float1 := 12312.123
// strt4 := strconv.FormatFloat(float1, 'b', 4, 32)

// str5 := fmt.Sprint(float1)
// str6 := fmt.Sprintf("%.2f", float1)
