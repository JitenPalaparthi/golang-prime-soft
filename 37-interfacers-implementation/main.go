package main

import (
	"fmt"
	"shapes/shape"
	"shapes/shape/circle"
	rectangle "shapes/shape/rect"
	"shapes/shape/square"
)

func main() {
	var shapes []shape.IShape
	shapes = append(shapes, rectangle.New(10.2, 13.5), rectangle.New(14.5, 16.78), square.New(12.4), square.New(123.23))
	shapes = append(shapes, circle.New(12.34), circle.New(2.23), rectangle.New(12.45, 23))
	for _, v := range shapes {
		Shape(v) // dynamic dispatch --> runtime polymorphism
		// when Shape calles v, v is an object of the type IShape so ,
		// every interface in go has a table called Interface Definition Table (IDT)
		// the same thing in c++, rust --> uses VTable
	}
	// what interface contain? Ptr to the actual data, pointer to the type
}

func Shape(iShape shape.IShape) {
	fmt.Println("Area of ", iShape.What(), "is", iShape.Area())
	fmt.Println("Perimeter of ", iShape.What(), "is", iShape.Perimeter())
	fmt.Println("---------")
}
