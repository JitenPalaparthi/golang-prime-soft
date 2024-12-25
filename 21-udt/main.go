package main

import "fmt"

func main() {

	var r1 Rect // not a pointer
	r1.L = 43.5
	r1.B = 39.0

	// a1 := Area(r1)
	// p1 := Perimeter(r1)
	//a1 := (&r1).Area() // no need of this
	a1 := r1.Area()
	p1 := r1.Perimeter()

	fmt.Println("Area a1 of Rect r1:", a1)
	fmt.Println("Perimeter p1 of Rect r1:", p1)
	fmt.Println(r1.A)
	fmt.Println(r1.P)

	r2 := Rect{L: 12.2, B: 3.4}
	r3 := &Rect{L: 12.2, B: 3.4}
	fmt.Println(r2, r3)

	r4 := new(Rect)
	r4.L = 12.4
	r4.B = 23.23
	a4 := r4.Area()
	p4 := r4.Perimeter()

	fmt.Println("Area a4 of Rect r4:", a4)
	fmt.Println("Perimeter p4 of Rect r4:", p4)

	r5 := New(12.34, 43.3) // similar to constructor

	a5 := r5.Area()
	p5 := r5.Perimeter()
	fmt.Println("Area a5 of Rect r5:", a5)
	fmt.Println("Perimeter p5 of Rect r5:", p5)

}

// type Person struct {
// 	Name    string
// 	Email   string
// 	Contact string
// }

type Rect struct {
	L, B, A, P float32
	// here cannot create or add methods
}

// type Square struct {
// 	S float32
// }

func New(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func NewDefault() *Rect {
	return &Rect{L: 1, B: 1}
}

// methods of a rect
func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

// functions

func Area(r Rect) float32 {
	return r.L * r.B
}

func Perimeter(r Rect) float32 {
	return 2 * (r.L + r.B)
}

// 1. UDT
