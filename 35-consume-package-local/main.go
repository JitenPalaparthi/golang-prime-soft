package main

import (
	"fmt"

	"github.com/jitenpalaparthi/prime-soft-shapes/shape"
	rectangle "github.com/jitenpalaparthi/prime-soft-shapes/shape/rect"
	sh "github.com/jitenpalaparthi/prime-soft-shapes/shape/square"
)

func main() {
	shape.What()
	r1 := rectangle.New(10, 20)
	a1 := r1.Area()
	p1 := r1.Perimeter()
	fmt.Println("Area of Rect r1", a1, "Perimeter of Rect r1", p1)

	var s1 sh.Square = sh.New(25.5)

	a2 := s1.Area()
	p2 := s1.Perimeter()
	fmt.Println("Area of Square s1", a2, "Perimeter of Square s1", p2)

}

// there is no public private in Go
// there is only exporeted or unexported
// restricted or unrestricted fields or types etc..
// any type , fileds,methods,functions inside a type or variable starts with Uppercase is exportable(unrestricted or similar to public)
// any type,field,methods, functions or a varialbe statrs with lowerCas (camelCase)is unexportable(restricted which is similar to private)
