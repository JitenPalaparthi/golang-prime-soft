package rectangle

var Global int // can be exported
var local int  //cannot be exported

func New(l, b float32) *Rect {
	return &Rect{l: l, b: b}
}

type Rect struct { // exported
	l, b float32 // unexported
}

func (r *Rect) what() string { // unexporeted
	return "Rect"
}

type t struct { // cannot be accessible in main as it startes with lowerCase
	L, B float32
}
