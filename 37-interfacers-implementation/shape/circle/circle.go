package circle

const PI = 3.14

func New(r float32) Circle {
	return Circle(r)
}

type Circle float32

func (c Circle) Area() float32 {
	return PI * float32(c*c)
}

func (c Circle) Perimeter() float32 {
	return 2 * PI * float32(c)
}

func (c Circle) What() string {
	return "Circle"
}
