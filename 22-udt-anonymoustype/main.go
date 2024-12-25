package main

import "fmt"

func main() {

	cc1 := ColourCode{int: 123, string: "red"}
	//fmt.Println(cc1)
	cc1.Display()

	cc2 := New(101, "Blue", "Some Blue")
	cc2.Display()

}

func New(i int, s string, ms string) *ColourCode {
	return &ColourCode{int: i, string: s, mystring: ms}
}

//type mystring = string

type ColourCode struct {
	int
	string
	mystring string
}

func (cc *ColourCode) Display() {
	fmt.Println("Code:", cc.int)
	fmt.Println("Colour:", cc.string)
	fmt.Println("Another Code:", cc.mystring)
}
