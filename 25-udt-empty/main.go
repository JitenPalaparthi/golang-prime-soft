package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var s1 S1

	a1 := s1.Area(10, 12)

	var s2 S2

	a2 := s2.Area(12.2)

	a3 := new(S1).Area(123.2, 12.23)

	fmt.Println(a1, a2, a3)

	var addr1 struct{ City string }
	addr1.City = "Hyd"

	var person1 struct{ Name string } = struct{ Name string }{Name: "Jiten"}
	var person2 Person = Person{Name: "Jiten"}

	var e1 struct{} // empty struct variable
	fmt.Println(e1)

	var e2 struct{} = struct{}{}

	fmt.Println(person1, person2, e2, "Size of e2", unsafe.Sizeof(e2), "Size of s1", unsafe.Sizeof(s1))

}

type Person struct {
	Name string
}

type S1 struct{} // empty struct
type S2 struct{} // empty struct

func (s *S1) Area(l, b float32) float32 {
	return l * b
}

func (s *S2) Area(side float32) float32 {
	return side * side
}
